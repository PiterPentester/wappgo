package utils

import (
	"net/http"
	"testing"
)

func TestInternalServerError(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InternalServerError(tt.args.w)
		})
	}
}
