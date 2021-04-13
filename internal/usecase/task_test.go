package usecase

import (
	"context"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
	"reflect"
	"testing"
)

func Test_uc_List(t1 *testing.T) {
	type fields struct {
		task storages.Task
	}
	type args struct {
		ctx       context.Context
		id        string
		createdAt string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &uc{
				task: tt.fields.task,
			}
			got, err := t.List(tt.args.ctx, tt.args.id, tt.args.createdAt)
			if (err != nil) != tt.wantErr {
				t1.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uc_Add(t1 *testing.T) {
	type fields struct {
		task storages.Task
	}
	type args struct {
		ctx  context.Context
		id   string
		date string
		task *entities.Task
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &uc{
				task: tt.fields.task,
			}
			if err := t.Add(tt.args.ctx, tt.args.id, tt.args.date, tt.args.task); (err != nil) != tt.wantErr {
				t1.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
