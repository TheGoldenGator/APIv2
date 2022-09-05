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
func (r *queryResolver) Member(ctx context.Context, search *string) (*model.Member, error) {
	panic(fmt.Errorf("not implemented: Member - member"))
}

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context, limitArg *int, pageArg *int, sort *model.MemberSort) (*model.MemberConnection, error) {
	panic(fmt.Errorf("not implemented: Members - members"))
}

// Streams is the resolver for the streams field.
func (r *queryResolver) Streams(ctx context.Context, limitArg *int, pageArg *int, status *model.StreamStatus) (*model.StreamConnection, error) {
	panic(fmt.Errorf("not implemented: Streams - streams"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
