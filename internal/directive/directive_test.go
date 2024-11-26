package directive_test

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/uudashr/iface/internal/directive"
)

func TestParseIgnore(t *testing.T) {
	testCases := map[string]struct {
		doc *ast.CommentGroup
		dir *directive.Ignore
	}{
		"no comment": {},
		"has directive": {
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "//iface:ignore",
					},
				},
			},
			dir: &directive.Ignore{},
		},
		"no directive": {
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "// just simple comment",
					},
				},
			},
			dir: nil,
		},
		"directive with one name": {
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "//iface:ignore=unused",
					},
				},
			},
			dir: &directive.Ignore{
				Names: []string{"unused"},
			},
		},
		"directive with two name": {
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "//iface:ignore=unused,identical",
					},
				},
			},
			dir: &directive.Ignore{
				Names: []string{"unused", "identical"},
			},
		},
		"directive with weird assignment": {
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "//iface:ignore-asd=unused",
					},
				},
			},
			dir: nil,
		},
		"directive empty val": {
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "//iface:ignore=",
					},
				},
			},
			dir: &directive.Ignore{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := directive.ParseIgnore(tc.doc)
			if !reflect.DeepEqual(got, tc.dir) {
				t.Errorf("ParseIgnore(%v) = %v, want %v", tc.doc, got, tc.dir)
			}
		})
	}
}

func TestIgnore_ShouldIgnore(t *testing.T) {
	testCases := map[string]struct {
		analyzerName string
		dir          *directive.Ignore
		expectIgnore bool
	}{
		"no directive - unused": {
			analyzerName: "unused",
		},
		"no directive - identical": {
			analyzerName: "unused",
		},
		"ignore all - unused": {
			analyzerName: "unused",
			dir:          &directive.Ignore{},
			expectIgnore: true,
		},
		"ignore all - identical": {
			analyzerName: "unused",
			dir:          &directive.Ignore{},
			expectIgnore: true,
		},
		"ignore unused - unused": {
			analyzerName: "unused",
			dir: &directive.Ignore{
				Names: []string{"unused"},
			},
			expectIgnore: true,
		},
		"ignore identical - identical": {
			analyzerName: "identical",
			dir: &directive.Ignore{
				Names: []string{"identical"},
			},
			expectIgnore: true,
		},
		"ignore identical - unused": {
			analyzerName: "unused",
			dir: &directive.Ignore{
				Names: []string{"identical"},
			},
		},
		"ignore unused,identical - unused": {
			analyzerName: "unused",
			dir: &directive.Ignore{
				Names: []string{"unused", "identical"},
			},
			expectIgnore: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			dir := tc.dir
			if got, want := dir != nil && dir.ShouldIgnore(tc.analyzerName), tc.expectIgnore; got != want {
				t.Errorf("dir.ShouldIgnore(%v, %q) = %v, want %v", tc.dir, tc.analyzerName, got, want)
			}
		})
	}
}
