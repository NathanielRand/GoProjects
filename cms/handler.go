package cms

import (
	"net/http"
	"strings"
	"time"
)

// ServePage function serves a page page based on the route matched
// This will match any URL beginning with /page
func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")

	if path == "" {
		pages, err := GetPages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Tmpl.ExecuteTemplate(w, "pages", pages)
		return
	}

	page, err := GetPage(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Tmpl.ExecuteTemplate(w, "page", page)
}

// ServePost function serves a post
func ServePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/post/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Post{
		Title:   strings.ToTitle(path),
		Content: "Here is my page",
		Comments: []*Comment{
			&Comment{
				Author:        "Nathaniel Rand",
				Comment:       "Looking good",
				DatePublished: time.Now(),
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "post", p)
}

// HandleNew function handles preview logic
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		// HTML input value "page"
		if contentType == "page" {
			p := &Page{
				Title:   title,
				Content: content,
			}
			_, err := CreatePage(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			Tmpl.ExecuteTemplate(w, "page", p)
			return
		}

		// HTML input value "post"
		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Projects CMS",
		Content: "Welcome Home!",
		Posts: []*Post{
			&Post{
				Title:         "Hello World",
				Content:       "Welcome to our world!",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "A post with comments",
				Content:       "Maybe this post will attract attention.",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "John Doe",
						Comment:       "Just a random comment",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}
