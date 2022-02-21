package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/umerm-work/hatch-assignment/graph/generated"
	"github.com/umerm-work/hatch-assignment/graph/model"
	"github.com/umerm-work/hatch-assignment/presentation"
)

func (r *mutationResolver) AddTodo(ctx context.Context, input model.CreateTodo) (*model.Todo, error) {
	in := presentation.ApiToModelCreateTodo(input)
	err := r.storage.Create(in)
	if err != nil {
		return nil, err
	}
	return presentation.ModelToApiTodo(in), nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	in := presentation.ApiToModelUpdateTodo(input)
	err := r.storage.Update(in)
	if err != nil {
		return nil, err
	}
	return presentation.ModelToApiTodo(in), nil
}

func (r *mutationResolver) RemoveTodo(ctx context.Context, input model.RemoveTodo) (*model.Todo, error) {
	err := r.storage.Delete(int64(input.ID))
	if err != nil {
		return nil, err
	}
	return &model.Todo{ID: input.ID}, nil
}

func (r *queryResolver) Todos(ctx context.Context, sortProperty *model.SortOptions, orderOption *model.OrderOptions) ([]*model.Todo, error) {
	sort := "created_at"
	order := "desc"
	if sortProperty != nil && *sortProperty == model.SortOptionsPriority {
		sort = model.SortOptionsPriority.String()
	}
	if orderOption != nil && *orderOption == model.OrderOptionsAsc {
		order = model.OrderOptionsAsc.String()
	}
	list, err := r.storage.List(sort, order)
	if err != nil {
		return nil, err
	}

	return presentation.ModelToApiTodos(list), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
