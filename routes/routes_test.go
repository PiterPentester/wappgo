package routes

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name string
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexGetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexGetHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_indexPostHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexPostHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_userGetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userGetHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_loginGetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginGetHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_loginPostHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginPostHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_logoutGetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logoutGetHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_registerGetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerGetHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_registerPostHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registerPostHandler(tt.args.w, tt.args.r)
		})
	}
}
