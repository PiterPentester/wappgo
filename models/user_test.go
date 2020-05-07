package models

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		username string
		hash     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.username, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetId(t *testing.T) {
	type fields struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				id: tt.fields.id,
			}
			got, err := user.GetId()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsername(t *testing.T) {
	type fields struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				id: tt.fields.id,
			}
			got, err := user.GetUsername()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetHash(t *testing.T) {
	type fields struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				id: tt.fields.id,
			}
			got, err := user.GetHash()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Authenticate(t *testing.T) {
	type fields struct {
		id int64
	}
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				id: tt.fields.id,
			}
			if err := user.Authenticate(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("User.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthenticateUser(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthenticateUser(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterUser(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterUser(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
