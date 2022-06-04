package markdown

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type ASTTransformer struct{}

func (g *ASTTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	_ = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch v := n.(type) {
		case *ast.Link:
			link := v.Destination
			if bytes.HasPrefix(link, []byte("http")) {
				v.SetAttributeString("target", []byte("_blank"))
			}
		}

		return ast.WalkContinue, nil
	})
}
