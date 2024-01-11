// Package markdown takes markdown and translates it into html
package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	mds := strings.Split(markdown, "\n")
	lst := false
	for i := range mds {
		mds[i] = regexp.MustCompile(`__(.*?)__`).ReplaceAllString(mds[i], `<strong>$1</strong>`)
		mds[i] = regexp.MustCompile(`_(.*?)_`).ReplaceAllString(mds[i], `<em>$1</em>`)
		if strings.HasPrefix(mds[i], "* ") {
			mds[i] = strings.Replace(mds[i], "* ", "<li>", 1) + "</li>"
			if !lst {
				mds[i] = "<ul>" + mds[i]
				lst = true
			}
			if (i < len(mds)-1 && !strings.HasPrefix(mds[i+1], "* ")) || i == len(mds)-1 {
				mds[i] = mds[i] + "</ul>"
				lst = false

			}
			continue
		}
		hd := regexp.MustCompile("^(#{1,6})( )").FindStringIndex(mds[i])
		if len(hd) > 0 {
			mds[i] = fmt.Sprintf("<h%d>", hd[1]-1) + mds[i][hd[1]:] + fmt.Sprintf("</h%d>", hd[1]-1)
			continue
		}
		mds[i] = "<p>" + mds[i] + "</p>"
	}
	html := strings.Join(mds, "")
	return html
}
