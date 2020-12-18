package graph

import "github.com/coreuml/go-graphql-server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	questions []*model.Question
	choices   []*model.Choice
}
