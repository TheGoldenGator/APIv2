package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/gql/graph/generated"
	"github.com/thegoldengator/APIv2/internal/gql/graph/model"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

// Member is the resolver for the member field.
func (r *streamResolver) Member(ctx context.Context, obj *model.Stream) (*model.Member, error) {
	// Fetch from database
	var member *structures.Member
	if err := database.Mongo.Members.FindOne(ctx, bson.M{"twitch_id": obj.TwitchID}).Decode(&member); err != nil {
		return nil, err
	}

	return &model.Member{
		ID:          string(member.ID.Hex()),
		TwitchID:    member.TwitchID,
		Login:       member.Login,
		DisplayName: member.DisplayName,
		Color:       member.Color,
		Pfp:         member.Pfp,
		Links:       (*model.MemberLink)(member.Links),
	}, nil
}

// Stream returns generated.StreamResolver implementation.
func (r *Resolver) Stream() generated.StreamResolver { return &streamResolver{r} }

type streamResolver struct{ *Resolver }
