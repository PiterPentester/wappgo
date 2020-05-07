package utils

import "testing"

func Test_getEnv(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnv(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("getEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
