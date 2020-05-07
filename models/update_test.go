package models

import (
	"reflect"
	"testing"
)

func TestNewUpdate(t *testing.T) {
	type args struct {
		userID int64
		body   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Update
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUpdate(tt.args.userID, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate_GetBody(t *testing.T) {
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
			update := &Update{
				id: tt.fields.id,
			}
			got, err := update.GetBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("Update.GetBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Update.GetBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate_GetUser(t *testing.T) {
	type fields struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			update := &Update{
				id: tt.fields.id,
			}
			got, err := update.GetUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("Update.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryUpdates(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Update
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := queryUpdates(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryUpdates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllUpdates(t *testing.T) {
	tests := []struct {
		name    string
		want    []*Update
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllUpdates()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUpdates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUpdates(t *testing.T) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*Update
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUpdates(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpdates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostUpdate(t *testing.T) {
	type args struct {
		userID int64
		body   string
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
			if err := PostUpdate(tt.args.userID, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("PostUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
