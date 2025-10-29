package wutl

import (
	"embed"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

type SkipFn func() bool

// ServeEmbedded returns a [HandlerFn] that serves embedde files from fileSys, root determines the root dir. You can pass optional
// skiprules if you want specific files to be handled differently
func ServeEmbedded(fileSys embed.FS, root string, skipsrules ...SkipFn) HandlerFn {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ext := filepath.Ext(r.URL.Path)
			if ext == "" {
				next.ServeHTTP(w, r)
				return
			}

			path := filepath.Join(root, strings.TrimPrefix(r.URL.Path, "/"))
			f, e := fileSys.Open(path)
			if e != nil {
				http.NotFound(w, r)
				return
			}
			defer f.Close()

			stats, _ := f.Stat()
			seeker, ok := f.(io.ReadSeeker)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.ServeContent(w, r, r.URL.Path, stats.ModTime(), seeker)
		})
	}
}
