package graph

import (
	"PostCommentService/Internal/database"
)

type Resolver struct {
	store database.Store
}

func NewResolver(useMemory bool) *Resolver {
	store := database.NewStore(useMemory)
	return &Resolver{store: store}
}
