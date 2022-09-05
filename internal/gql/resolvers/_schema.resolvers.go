package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/thegoldengator/APIv2/internal/gql/generated"
	"github.com/thegoldengator/APIv2/internal/gql/graph/model"
)

// Member is the resolver for the member field.
func (r *queryResolver) Member(ctx context.Context, search *string) (*model.Member, error) {
	if search == nil {
		return nil, errors.New("please provide a user_id, login, or display name")
	}

	found, err := cache.Members.Search(ctx, *search)
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
	}, nil
}

// Members is the resolver for the members field.
func (r *queryResolver) Members(ctx context.Context, sort *model.MemberSort, first *int, after *string) (*model.MemberConnection, error) {
	var decodedCursor string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}

	var edges = make([]*model.MemberEdge, 0)
	count := 0
	currentPage := false

	if decodedCursor == "" {
		currentPage = true
	}
	hasNextPage := false

	var members []*structures.Member
	var err error
	if decodedCursor == "" {
		members, err = cache.Members.GetByRange(ctx, 1, *first)
		if err != nil {
			return nil, err
		}
	} else {
		members, err = cache.Members.GetAfterID(ctx, "az", decodedCursor, *first)
		if err != nil {
			return nil, err
		}
	}

	edges = nil
	for _, s := range members {
		newNode := &model.Member{
			ID:          s.ID.Hex(),
			TwitchID:    s.TwitchID,
			Login:       s.Login,
			DisplayName: s.DisplayName,
			Pfp:         s.Pfp,
			Links:       (*model.MemberLink)(s.Links),
		}

		edges = append(edges, &model.MemberEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(s.TwitchID)),
			Node:   newNode,
		})
	}

	for i, edg := range edges {
		if edg.Node.TwitchID == decodedCursor {
			currentPage = true
		}

		if currentPage && count < *first {
			edges[count] = &model.MemberEdge{
				Cursor: base64.StdEncoding.EncodeToString([]byte(edg.Node.TwitchID)),
				Node:   edg.Node,
			}
			count++
		}

		if count == *first && i < len(edges) {
			hasNextPage = true
		}
	}

	firstEdge := edges[0].Cursor
	lastEdge := edges[len(edges)-1].Cursor

	pageInfo := model.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(firstEdge)),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(lastEdge)),
		HasNextPage: &hasNextPage,
	}

	return &model.MemberConnection{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

// Cache by status, if user wants all streams it'll merge the two.
func (r *queryResolver) Streams(ctx context.Context, status *model.StreamStatus, first *int, after *string) (*model.StreamConnection, error) {
	var decodedCursor string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}

	var edges = make([]*model.StreamEdge, 0)
	count := 0
	currentPage := false

	// if no cursor start from top
	if decodedCursor == "" {
		currentPage = true
	}
	hasNextPage := false

	var streams []*structures.Stream
	var err error
	if decodedCursor == "" {
		streams, err = cache.Streams.GetByRange(status.String(), 0, *first)
		if err != nil {
			return nil, err
		}
	} else {
		streams, err = cache.Streams.GetAfterID(ctx, status.String(), decodedCursor, *first)
		if err != nil {
			return nil, err
		}
	}

	edges = nil
	for _, s := range streams {
		newNode := &model.Stream{
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
		}

		member, err := r.Resolver.Stream().Member(ctx, newNode)
		if err != nil {
			return nil, err
		}

		newNode.Member = member

		edges = append(edges, &model.StreamEdge{
			Cursor: base64.StdEncoding.EncodeToString([]byte(s.TwitchID)),
			Node:   newNode,
		})
	}

	for i, se := range edges {
		if se.Node.TwitchID == decodedCursor {
			currentPage = true
		}

		if currentPage && count < *first {
			edges[count] = &model.StreamEdge{
				Cursor: base64.StdEncoding.EncodeToString([]byte(se.Node.TwitchID)),
				Node:   se.Node,
			}
			count++
		}

		if count == *first && i < len(edges) {
			hasNextPage = true
		}
	}

	firstEdge := edges[0].Cursor
	lastEdge := edges[len(edges)-1].Cursor

	pageInfo := model.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(firstEdge)),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(lastEdge)),
		HasNextPage: &hasNextPage,
	}

	return &model.StreamConnection{
		Edges:    edges,
		PageInfo: &pageInfo,
	}, nil
}

// Stats is the resolver for the stats field.
func (r *queryResolver) Stats(ctx context.Context, search *string) ([]*model.StatEntry, error) {
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
