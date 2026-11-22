package mcp

import (
	"bufio"
	"fmt"
	"io/fs"
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
func (idx *DocIndexer) Get(path string, locale string) (*DocPage, error) {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if locale == "" {
		locale = "en"
	}
	cleanTarget := stripNumericPrefixes(normalizePath(path))

	for _, page := range idx.pages {
		if page.Locale != locale && locale != "all" {
			continue
		}
		if normalizePath(page.Path) == normalizePath(path) ||
			page.CleanPath == cleanTarget ||
			strings.HasSuffix(normalizePath(page.Path), normalizePath(path)) ||
			strings.HasSuffix(page.CleanPath, cleanTarget) {
			return &page, nil
		}
	}

	return nil, fmt.Errorf("documentation page not found: %s (locale: %s)", path, locale)
}

// List returns all documentation pages for a given locale.
func (idx *DocIndexer) List(locale string) []DocPage {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if locale == "" {
		locale = "en"
	}

	var out []DocPage
	for _, page := range idx.pages {
		if page.Locale == locale || locale == "all" {
			// Return a shallow copy with truncated content for listing
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
	// Normalize path: e.g. "content/en/docs/02-getting-started/01-quickstart.md"
	rel := fullPath
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

	canonicalURL := "https://krewire.com/docs/" + cleanPath
	if locale == "id" {
		canonicalURL = "https://krewire.com/id/docs/" + cleanPath
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

func normalizePath(p string) string {
	p = strings.Trim(p, "/")
	p = strings.TrimPrefix(p, "content/")
	p = strings.TrimPrefix(p, "en/docs/")
	p = strings.TrimPrefix(p, "id/docs/")
	p = strings.TrimPrefix(p, "docs/")
	p = strings.TrimSuffix(p, ".md")
	p = strings.TrimSuffix(p, ".html")
	return filepath.ToSlash(p)
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
