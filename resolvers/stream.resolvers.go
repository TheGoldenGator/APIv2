package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/thegoldengator/APIv2/internal/gql/graph/generated"
	"github.com/thegoldengator/APIv2/internal/gql/graph/model"
)

// Member is the resolver for the member field.
func (r *streamResolver) Member(ctx context.Context, obj *model.Stream) (*model.Member, error) {
	panic(fmt.Errorf("not implemented: Member - member"))
}

// Stream returns generated.StreamResolver implementation.
func (r *Resolver) Stream() generated.StreamResolver { return &streamResolver{r} }

type streamResolver struct{ *Resolver }
