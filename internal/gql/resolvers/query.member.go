package resolvers

import (
	"context"
	"errors"

	"github.com/thegoldengator/APIv2/internal/gql/graph/model"
)

func (r *queryResolver) Member(ctx context.Context, search *string) (*model.Member, error) {
	if search == nil {
		return nil, errors.New("please provide a user_id, login, or display name")
	}

	/* found, err := cache.Members.Search(ctx, *search)
	if err != nil {
		return nil, err
	}

	return &model.Member{
		ID:          found.ID.Hex(),
		TwitchID:    found.TwitchID,
		Login:       found.Login,
		DisplayName: found.DisplayName,
		Pfp:         found.Pfp,
		Links:       (*model.MemberLink)(found.Links),
	}, nil */
	return nil, nil
}
