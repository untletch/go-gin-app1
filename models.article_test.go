package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	if !slices.Equal(alist, articleList) {
		t.Fail()
	}
}
