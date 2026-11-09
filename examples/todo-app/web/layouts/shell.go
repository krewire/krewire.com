// Package layouts defines the application page shells (FRK-STR-006).
package layouts

import "github.com/krewire/krewire/pkg/web/ssg"

const shellBody = "<!DOCTYPE html>\n" +
"<html lang=\"en\">\n" +
"<head>\n" +
"<meta charset=\"utf-8\">\n" +
"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n" +
"<meta name=\"color-scheme\" content=\"light dark\">\n" +
"<title>{{.Title}}</title>\n" +
"<link rel=\"stylesheet\" href=\"/assets/ui.css\">\n" +
"<link rel=\"stylesheet\" href=\"/assets/theme.css\">\n" +
"{{themeHead}}\n" +
"</head>\n" +
"<body>\n" +
"<header class=\"topbar\"><strong>Krewire App</strong><span class=\"spacer\"></span>{{themeButton}}</header>\n" +
"<main>{{.Content}}</main>\n" +
"<footer class=\"footer\">Served by a single krewire binary.</footer>\n" +
"</body>\n" +
"</html>"

// Shell is the default application layout: topbar with theme toggle, content,
// footer.
func Shell() ssg.Layout {
	return ssg.Layout{
		Name:  "Shell",
		Body:  shellBody,
		Style: ".topbar{display:flex;align-items:center;gap:1rem}.spacer{flex:1}",
	}
}
