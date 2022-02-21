package presentation

import (
	"github.com/umerm-work/hatch-assignment/data"
	"github.com/umerm-work/hatch-assignment/graph/model"
	"time"
)

func ModelToApiTodo(in data.Todo) *model.Todo {
	return &model.Todo{
		ID:          int(in.ID),
		Title:       in.Title,
		Description: in.Description,
		Priority:    in.Priority,
		CreatedAt:   ConvertTimestamp(&in.CreatedAt),
	}
}

func ModelToApiTodos(in []data.Todo) []*model.Todo {
	todos := make([]*model.Todo, 0, len(in))
	for _, todo := range in {
		todos = append(todos, ModelToApiTodo(todo))
	}
	return todos
}

//func ApiToModelTodo(in model.Todo) data.Todo {
//	return data.Todo{
//		ID:          int64(in.ID),
//		Title:       in.Title,
//		Description: in.Description,
//		Priority:    in.Priority,
//	}
//}
func ApiToModelCreateTodo(in model.CreateTodo) data.Todo {
	return data.Todo{
		Title:       in.Title,
		Description: in.Description,
		Priority:    in.Priority,
	}
}
func ApiToModelUpdateTodo(in model.UpdateTodo) data.Todo {
	return data.Todo{
		ID:          int64(in.ID),
		Title:       in.Title,
		Description: in.Description,
		Priority:    in.Priority,
	}
}
func ConvertTimestamp(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02")
	return &s
}
