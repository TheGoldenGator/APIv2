package twitch

import (
	"context"
	"net/http"

	"github.com/thegoldengator/APIv2/pkg/structures"
	"go.mongodb.org/mongo-driver/bson"
)

// Searches streams using pagination
func (t Twitch) SearchStreams(ctx context.Context, filter bson.M, opts ...structures.UserSearchOptions) ([]structures.Stream, int, error) {

	return nil, http.StatusOK, nil
}
