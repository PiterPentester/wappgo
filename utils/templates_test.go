package utils

import (
	"net/http"
	"testing"
)

func TestLoadTemplates(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadTemplates(tt.args.pattern)
		})
	}
}

func TestExecuteTemplate(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		tmpl string
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExecuteTemplate(tt.args.w, tt.args.tmpl, tt.args.data)
		})
	}
}
