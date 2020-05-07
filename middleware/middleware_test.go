package middleware

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthRequired(t *testing.T) {
	type args struct {
		handler http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthRequired(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
