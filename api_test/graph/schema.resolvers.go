package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api_test/graph/generated"
	"api_test/graph/model"
	"api_test/graph/models"
	"context"
	"fmt"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id int) (*models.Todo, error) {
	db := r.Resolver.DB
	var todo models.Todo
	err := db.Debug().First(&todo, id).Error

	return &todo, err
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*model.User, error) {
	db := r.Resolver.DB
	var user model.User
	err := db.Debug().First(&user, obj.UserID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
