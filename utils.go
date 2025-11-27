package wutl

import (
	"errors"
	"path/filepath"
	"strings"
)

var ErrDirOutOfBounds = errors.New("outside of directory bounds")

// sanitize_path makes sure input can't access anything outside of root
func sanitize_path(root, path string) (string, error) {
	absolutepath, e := filepath.Abs(root)
	if e != nil {
		return "", e
	}

	max_surface := len(strings.Split(absolutepath, "/"))
	if strings.HasPrefix(path, root) {
		path = strings.Replace(path, root, "", 1)
	}
	path = filepath.Join(absolutepath, strings.TrimPrefix(path, "/"))

	parts := strings.Split(path, "/")
	path_depth := len(parts)

	for _, p := range parts {
		if p == ".." {
			path_depth -= 1
		}
	}

	if path_depth < max_surface {
		return "", ErrDirOutOfBounds
	}

	return path, nil
}

