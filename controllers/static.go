package controllers

import (
	"net/http"

	"github.com/lemon57/share-moment/views"
)

type Static struct {
	Template views.Template
}

func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
