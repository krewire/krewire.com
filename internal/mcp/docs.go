package mcp

import (
	"bufio"
	"fmt"
	"io/fs"
	"net/url"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// DocPage represents an indexed documentation document.
type DocPage struct {
	Path        string   `json:"path"`       // e.g. "02-getting-started/01-quickstart"
	CleanPath   string   `json:"clean_path"` // e.g. "getting-started/quickstart"
	Locale      string   `json:"locale"`     // "en" or "id"
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
	Content     string   `json:"content"`
	URL         string   `json:"url"` // e.g. "https://krewire.com/docs/getting-started/quickstart"
}

// SearchResult contains a search match with scoring.
type SearchResult struct {
	Page    DocPage `json:"page"`
	Score   int     `json:"score"`
	Excerpt string  `json:"excerpt"`
}

// DocIndexer provides in-memory indexing and full-text search over Krewire documentation.
type DocIndexer struct {
	mu    sync.RWMutex
	pages []DocPage
}

// NewDocIndexer parses and indexes markdown documents from the provided filesystem.
func NewDocIndexer(contentFS fs.FS) (*DocIndexer, error) {
	idx := &DocIndexer{}
	if contentFS == nil {
		return idx, nil
	}

	err := fs.WalkDir(contentFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		data, err := fs.ReadFile(contentFS, path)
		if err != nil {
			return err
		}

		page := parseDocPage(path, string(data))
		idx.pages = append(idx.pages, page)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("indexing docs: %w", err)
	}

	return idx, nil
}

// Count returns the total number of indexed documentation pages.
func (idx *DocIndexer) Count() int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()
	return len(idx.pages)
}

// Search finds documentation pages matching the given query string.
func (idx *DocIndexer) Search(query string, locale string, limit int) []SearchResult {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if limit <= 0 {
		limit = 5
	}
	if locale == "" {
		locale = "en"
	} else {
		locale = strings.ToLower(strings.TrimSpace(locale))
	}

	tokens := strings.Fields(strings.ToLower(query))
	if len(tokens) == 0 {
		return nil
	}

	var results []SearchResult
	for _, page := range idx.pages {
		if page.Locale != locale && locale != "all" {
			continue
		}

		score := 0
		titleLower := strings.ToLower(page.Title)
		descLower := strings.ToLower(page.Description)
		pathLower := strings.ToLower(page.Path)
		contentLower := strings.ToLower(page.Content)

		for _, token := range tokens {
			if strings.Contains(titleLower, token) {
				score += 15
			}
			if strings.Contains(descLower, token) {
				score += 8
			}
			if strings.Contains(pathLower, token) {
				score += 6
			}
			for _, tag := range page.Tags {
				if strings.Contains(strings.ToLower(tag), token) {
					score += 10
				}
			}
			if count := strings.Count(contentLower, token); count > 0 {
				score += count * 2
				if score > 100 {
					score = 100
				}
			}
		}

		if score > 0 {
			excerpt := generateExcerpt(page.Content, tokens)
			results = append(results, SearchResult{
				Page:    page,
				Score:   score,
				Excerpt: excerpt,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	if len(results) > limit {
		results = results[:limit]
	}

	return results
}

// Get retrieves a single documentation page by path and locale.
// Supports full paths, clean paths, .md/.html extensions, URLs, and directory names.
func (idx *DocIndexer) Get(rawPath string, locale string) (*DocPage, error) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	rawPath = strings.TrimSpace(rawPath)
	if rawPath == "" {
		return nil, fmt.Errorf("path argument is required. Use list_docs to view available paths")
	}

	normPath, detectedLocale := normalizeDocInputPath(rawPath)
	if locale == "" {
		if detectedLocale != "" {
			locale = detectedLocale
		} else {
			locale = "en"
		}
	} else {
		locale = strings.ToLower(strings.TrimSpace(locale))
	}

	cleanTarget := stripNumericPrefixes(normPath)

	// Phase 1: Exact matches by Path or CleanPath
	for _, page := range idx.pages {
		if page.Locale != locale && locale != "all" {
			continue
		}
		if page.Path == normPath ||
			page.CleanPath == normPath ||
			page.Path == cleanTarget ||
			page.CleanPath == cleanTarget {
			return &page, nil
		}
	}

	// Phase 2: Root index aliases
	if normPath == "index" || normPath == "00-index" || normPath == "" {
		for _, page := range idx.pages {
			if page.Locale != locale && locale != "all" {
				continue
			}
			if page.Path == "00-index" || page.CleanPath == "index" {
				return &page, nil
			}
		}
	}

	// Phase 3: Directory index aliases (e.g. "getting-started" -> "02-getting-started/00-index")
	for _, page := range idx.pages {
		if page.Locale != locale && locale != "all" {
			continue
		}
		if page.Path == normPath+"/00-index" ||
			page.Path == normPath+"/index" ||
			page.Path == cleanTarget+"/00-index" ||
			(cleanTarget != "" && page.CleanPath == cleanTarget) {
			return &page, nil
		}
	}

	// Phase 4: Suffix segment matches (e.g. "08-mcp-server" -> "02-getting-started/08-mcp-server")
	if normPath != "" {
		for _, page := range idx.pages {
			if page.Locale != locale && locale != "all" {
				continue
			}
			if strings.HasSuffix(page.Path, "/"+normPath) ||
				(cleanTarget != "" && cleanTarget != "index" && strings.HasSuffix(page.CleanPath, "/"+cleanTarget)) {
				return &page, nil
			}
		}
	}

	return nil, fmt.Errorf("documentation page not found: %q (locale: %q). Available paths can be discovered using list_docs", rawPath, locale)
}

// List returns all documentation pages for a given locale.
func (idx *DocIndexer) List(locale string) []DocPage {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if locale == "" {
		locale = "en"
	} else {
		locale = strings.ToLower(strings.TrimSpace(locale))
	}

	var out []DocPage
	for _, page := range idx.pages {
		if page.Locale == locale || locale == "all" {
			// Return a shallow copy with empty content for compact listing
			p := page
			p.Content = ""
			out = append(out, p)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Path < out[j].Path
	})

	return out
}

func parseDocPage(fullPath, raw string) DocPage {
	rel := filepath.ToSlash(fullPath)
	rel = strings.TrimPrefix(rel, "content/")
	rel = strings.TrimSuffix(rel, ".md")

	parts := strings.Split(rel, "/")
	locale := "en"
	if len(parts) > 0 && (parts[0] == "en" || parts[0] == "id") {
		locale = parts[0]
		parts = parts[1:]
	}
	if len(parts) > 0 && parts[0] == "docs" {
		parts = parts[1:]
	}

	subPath := strings.Join(parts, "/")
	cleanPath := stripNumericPrefixes(subPath)

	title := cleanTitleFromPath(parts)
	desc := ""
	var tags []string

	body := raw
	if strings.HasPrefix(raw, "---") {
		parts := strings.SplitN(raw, "---", 3)
		if len(parts) >= 3 {
			fm := parts[1]
			body = strings.TrimSpace(parts[2])

			scanner := bufio.NewScanner(strings.NewReader(fm))
			inTags := false
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if strings.HasPrefix(line, "title:") {
					title = cleanYAMLString(strings.TrimPrefix(line, "title:"))
					inTags = false
				} else if strings.HasPrefix(line, "description:") {
					desc = cleanYAMLString(strings.TrimPrefix(line, "description:"))
					inTags = false
				} else if strings.HasPrefix(line, "tags:") {
					inTags = true
				} else if inTags && strings.HasPrefix(line, "- ") {
					tags = append(tags, strings.TrimPrefix(line, "- "))
				} else if inTags && !strings.HasPrefix(line, "-") && strings.Contains(line, ":") {
					inTags = false
				}
			}
		}
	}

	if desc == "" {
		desc = fallbackDescription(body)
	}

	canonicalURL := "https://krewire.com/docs"
	if cleanPath != "" && cleanPath != "index" {
		canonicalURL = "https://krewire.com/docs/" + cleanPath
	}
	if locale == "id" {
		canonicalURL = "https://krewire.com/id/docs"
		if cleanPath != "" && cleanPath != "index" {
			canonicalURL = "https://krewire.com/id/docs/" + cleanPath
		}
	}

	return DocPage{
		Path:        subPath,
		CleanPath:   cleanPath,
		Locale:      locale,
		Title:       title,
		Description: desc,
		Tags:        tags,
		Content:     body,
		URL:         canonicalURL,
	}
}

func stripNumericPrefixes(path string) string {
	segments := strings.Split(path, "/")
	var clean []string
	for _, s := range segments {
		s = stripSinglePrefix(s)
		if s != "" && s != "index" && s != "00-index" {
			clean = append(clean, s)
		}
	}
	if len(clean) == 0 {
		return "index"
	}
	return strings.Join(clean, "/")
}

func stripSinglePrefix(s string) string {
	parts := strings.SplitN(s, "-", 2)
	if len(parts) == 2 {
		isNumeric := true
		for _, c := range parts[0] {
			if c < '0' || c > '9' {
				isNumeric = false
				break
			}
		}
		if isNumeric {
			return parts[1]
		}
	}
	return s
}

func cleanTitleFromPath(parts []string) string {
	if len(parts) == 0 {
		return "Documentation"
	}
	last := parts[len(parts)-1]
	last = stripSinglePrefix(last)
	last = strings.ReplaceAll(last, "-", " ")
	return strings.Title(last)
}

func cleanYAMLString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, `"'`)
	return s
}

func normalizeDocInputPath(p string) (normPath string, detectedLocale string) {
	p = strings.TrimSpace(p)

	// If URL, parse out the path
	if strings.Contains(p, "://") {
		if u, err := url.Parse(p); err == nil {
			p = u.Path
		}
	}

	// Remove query and fragment if present
	if idx := strings.IndexAny(p, "?#"); idx != -1 {
		p = p[:idx]
	}

	p = filepath.ToSlash(p)
	p = strings.Trim(p, "/")

	// Strip common container and repo root prefixes
	p = strings.TrimPrefix(p, "usr/share/nginx/html/")
	p = strings.TrimPrefix(p, "content/")

	// Detect and extract locale if encoded in path
	if strings.HasPrefix(p, "id/docs/") {
		detectedLocale = "id"
		p = strings.TrimPrefix(p, "id/docs/")
	} else if strings.HasPrefix(p, "en/docs/") {
		detectedLocale = "en"
		p = strings.TrimPrefix(p, "en/docs/")
	} else if strings.HasPrefix(p, "id/") {
		detectedLocale = "id"
		p = strings.TrimPrefix(p, "id/")
	} else if strings.HasPrefix(p, "en/") {
		detectedLocale = "en"
		p = strings.TrimPrefix(p, "en/")
	}

	p = strings.TrimPrefix(p, "docs/")
	p = strings.Trim(p, "/")
	p = strings.TrimSuffix(p, ".md")
	p = strings.TrimSuffix(p, ".html")

	return p, detectedLocale
}

func fallbackDescription(body string) string {
	lines := strings.Split(body, "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) > 20 && !strings.HasPrefix(l, "#") && !strings.HasPrefix(l, "---") && !strings.HasPrefix(l, "```") {
			if len(l) > 160 {
				return l[:157] + "..."
			}
			return l
		}
	}
	return "Documentation manuscript"
}

func generateExcerpt(content string, tokens []string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		lineClean := strings.TrimSpace(line)
		if len(lineClean) < 20 || strings.HasPrefix(lineClean, "#") || strings.HasPrefix(lineClean, "---") {
			continue
		}
		lineLower := strings.ToLower(lineClean)
		for _, token := range tokens {
			if strings.Contains(lineLower, token) {
				if len(lineClean) > 200 {
					return lineClean[:200] + "..."
				}
				return lineClean
			}
		}
	}

	for _, line := range lines {
		lineClean := strings.TrimSpace(line)
		if len(lineClean) > 30 && !strings.HasPrefix(lineClean, "#") {
			if len(lineClean) > 200 {
				return lineClean[:200] + "..."
			}
			return lineClean
		}
	}
	return ""
}
