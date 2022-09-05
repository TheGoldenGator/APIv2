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

const MAX_STREAMS = 100

// Cache by status, if user wants all streams it'll merge the two.
func (r *queryResolver) Streams(ctx context.Context, limitArg *int, pageArg *int, status *model.StreamStatus) (*model.StreamConnection, error) {
	limit := MAX_STREAMS
	if limitArg != nil {
		limit = *limitArg
	}

	if limit > MAX_STREAMS {
		limit = MAX_STREAMS
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

	var streams []structures.Stream
	var filter bson.M

	if status != nil {
		filter = bson.M{"status": status.String()}
	} else {
		filter = bson.M{}
	}

	paginatedData, err := paginate.New(database.Mongo.Stream).Context(ctx).Limit(int64(limit)).Page(int64(page)).Sort("viewers", -1).Filter(filter).Decode(&streams).Find()
	if err != nil {
		return nil, err
	}

	var modeledStreams []*model.Stream
	for _, s := range streams {
		modeledStreams = append(modeledStreams, &model.Stream{
			ID:        s.ID.Hex(),
			TwitchID:  s.TwitchID,
			Member:    nil,
			Status:    model.StreamStatus(s.Status),
			Title:     s.Title,
			GameID:    s.GameID,
			Game:      s.Game,
			Viewers:   s.Viewers,
			Thumbnail: s.Thumbnail,
			StartedAt: s.StartedAt,
		})
	}

	return &model.StreamConnection{
		Streams: modeledStreams,
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
