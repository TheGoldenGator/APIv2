package resolvers

import (
	"context"
	"errors"

	paginate "github.com/gobeam/mongo-go-pagination"
	"github.com/thegoldengator/APIv2/internal/database"
	"github.com/thegoldengator/APIv2/internal/gql/graph/model"
	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

const MAX_MEMBERS = 100

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context, limitArg *int, pageArg *int, sort *model.MemberSort) (*model.MemberConnection, error) {
	limit := MAX_MEMBERS
	if limitArg != nil {
		limit = *limitArg
	}

	if limit > MAX_MEMBERS {
		limit = MAX_MEMBERS
	} else if limit < 1 {
		return nil, errors.New("limit cannot be less than 1")
	}

	page := 1
	if pageArg != nil {
		page = *pageArg
	}

	if page < 1 {
		page = 1
	}

	var members []structures.Member
	filter := bson.M{}
	var sortNumber int
	if sort != nil {
		switch *sort {
		case model.MemberSortAz:
			sortNumber = 1
		case model.MemberSortZa:
			sortNumber = -1
		}
	}

	paginatedData, err := paginate.New(database.Mongo.Members).Context(ctx).Limit(int64(limit)).Page(int64(page)).Sort("login", sortNumber).Filter(filter).Decode(&members).Find()
	if err != nil {
		return nil, err
	}

	var modeledMembers []*model.Member
	for _, s := range members {
		modeledMembers = append(modeledMembers, &model.Member{
			ID:          s.ID.Hex(),
			TwitchID:    s.TwitchID,
			Login:       s.Login,
			DisplayName: s.DisplayName,
			Color:       s.Color,
			Pfp:         s.Pfp,
			Links:       (*model.MemberLink)(s.Links),
		})
	}

	return &model.MemberConnection{
		Members: modeledMembers,
		PageInfo: &model.PageInfo{
			Total:     paginatedData.Pagination.Total,
			Page:      paginatedData.Pagination.Page,
			PerPage:   paginatedData.Pagination.PerPage,
			Prev:      paginatedData.Pagination.Prev,
			Next:      paginatedData.Pagination.Next,
			TotalPage: paginatedData.Pagination.TotalPage,
		},
	}, nil
}
