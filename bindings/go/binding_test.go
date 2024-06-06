package tree_sitter_git_commit_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-git_commit"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_git_commit.Language())
	if language == nil {
		t.Errorf("Error loading GitCommit grammar")
	}
}
