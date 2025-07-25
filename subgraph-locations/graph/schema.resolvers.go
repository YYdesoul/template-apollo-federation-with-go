package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.75

import (
	"context"
	"fmt"

	"github.com/YYdesoul/template-apollo-federation-with-go/subgraph-locations/graph/model"
)

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented: Locations - locations"))
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, id string) (*model.Location, error) {
	return &model.Location{
		ID: id,
		Name: "sample name",
		Description: "sample description",
		Photo: "sample Photo",
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
type mutationResolver struct{ *Resolver }
*/
