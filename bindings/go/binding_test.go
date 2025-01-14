package tree_sitter_markdown_test

import (
	"testing"

	tree_sitter_markdown "github.com/ponyruntime/tree-sitter-markdown/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadBlockGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_markdown.Language())
	if language == nil {
		t.Errorf("Error loading Markdown block grammar")
	}
}