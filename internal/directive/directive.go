package directive

import (
	"go/ast"
	"slices"
	"strings"
)

// Ignore represent a special instruction embebded in the source code.
//
// The directive can be as simple as
//
//	//iface:ignore
//
// or consist of name
//
//	//iface:ignore=unused
//
// or multiple names
//
//	//iface:ignore=unused,identical
type Ignore struct {
	Names []string
}

// ParseIgnore parse the directive from the comments.
func ParseIgnore(doc *ast.CommentGroup) *Ignore {
	if doc == nil {
		return nil
	}

	for _, comment := range doc.List {
		if !strings.HasPrefix(comment.Text, "//iface:ignore") {
			continue
		}

		// parse the Names if exists
		if strings.Contains(comment.Text, "=") {
			parts := strings.Split(comment.Text, "=")
			if len(parts) != 2 {
				continue
			}

			names := strings.Split(parts[1], ",")
			if len(names) == 0 {
				continue
			}

			for i, name := range names {
				names[i] = strings.TrimSpace(name)
			}

			if len(names) > 0 {
				return &Ignore{Names: names}
			}
		}

		return &Ignore{}
	}

	return nil
}

func (i *Ignore) hasName(name string) bool {
	return slices.Contains(i.Names, name)
}

// ShouldIgnore return true if the name should be ignored.
func (i *Ignore) ShouldIgnore(name string) bool {
	if len(i.Names) == 0 {
		return true
	}

	return i.hasName(name)
}
