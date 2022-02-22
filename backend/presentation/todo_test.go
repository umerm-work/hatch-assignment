package presentation

import (
	"github.com/umerm-work/hatch-assignment/data"
	"github.com/umerm-work/hatch-assignment/graph/model"
	"log"
	"reflect"
	"testing"
	"time"
)

const (
	TestStringData = "abc"
	TestNumberData = 1
)

func TestApiToModelCreateTodo(t *testing.T) {
	tests := []struct {
		name string
		in   model.CreateTodo
		want data.Todo
	}{
		{
			name: "Success Test",
			in: model.CreateTodo{
				Title:       TestStringData,
				Description: TestStringData,
				Priority:    TestNumberData,
			},
			want: data.Todo{
				Title:       TestStringData,
				Description: TestStringData,
				Priority:    TestNumberData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApiToModelCreateTodo(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiToModelCreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiToModelUpdateTodo(t *testing.T) {
	tests := []struct {
		name string
		in   model.UpdateTodo
		want data.Todo
	}{
		{
			name: "Success Test",
			in: model.UpdateTodo{
				Title:       TestStringData,
				Description: TestStringData,
				Priority:    TestNumberData,
			},
			want: data.Todo{
				Title:       TestStringData,
				Description: TestStringData,
				Priority:    TestNumberData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApiToModelUpdateTodo(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiToModelUpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTimestamp(t *testing.T) {
	testTime := time.Date(2022, 1, 1, 11, 11, 11, 0, &time.Location{})
	testTimeStr := testTime.Format("2006-01-02")
	log.Println(testTime)
	log.Println(testTimeStr)
	tests := []struct {
		name string
		t    *time.Time
		want *string
	}{
		{
			name: "Success",
			t:    &testTime,
			want: &testTimeStr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertTimestamp(tt.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelToApiTodo(t *testing.T) {
	tests := []struct {
		name string
		in   data.Todo
		want *model.Todo
	}{
		{
			name: "Success Test",
			in: data.Todo{
				ID:          TestNumberData,
				Title:       TestStringData,
				Description: TestStringData,
				Priority:    TestNumberData,
			},
			want: &model.Todo{
				ID:          TestNumberData,
				Title:       TestStringData,
				Description: TestStringData,
				Priority:    TestNumberData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModelToApiTodo(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelToApiTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelToApiTodos(t *testing.T) {
	tests := []struct {
		name string
		in   []data.Todo
		want []*model.Todo
	}{
		{
			name: "Success Test",
			in: []data.Todo{
				{ID: TestNumberData,
					Title:       TestStringData,
					Description: TestStringData,
					Priority:    TestNumberData},
			},
			want: []*model.Todo{
				{
					ID:          TestNumberData,
					Title:       TestStringData,
					Description: TestStringData,
					Priority:    TestNumberData,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModelToApiTodos(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelToApiTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}
