package structures

import "go.mongodb.org/mongo-driver/bson"

type UserSearchOptions struct {
	Page  int
	Limit int
	Query string
	Sort  bson.M
}
