package cart

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Item structure
type Product struct {
	ID   primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}


// GenerateItemID adds item "_id" in db
func (it *Product) GenerateItemID() {
	it.ID = primitive.NewObjectID()
}

// Validate checks items request body
func (it *Product) Validate() error {
	if it.Type == "" {
		return errors.New("invalid product type")

	}
	if it.Data == nil {
		return errors.New("product data can not be nil")

	}
	return nil
}