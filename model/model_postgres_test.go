package model

import (
	"testing"

	"github.com/go-pg/pg/v10"
)

func Test_createSchema(t *testing.T) {
	type args struct {
		db *pg.DB
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
			if err := createSchema(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("createSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnectPGDatabase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "postgres://postgres:postgres@localhost:5434/togo?sslmode=disable",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectPGDatabase()
		})
	}
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize()
		})
	}
}

func TestClosePGConnection(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClosePGConnection()
		})
	}
}
