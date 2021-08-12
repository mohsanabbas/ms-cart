package cart

import (
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	expirationTime = 24
)

// Cart response structure
type Cart struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Expire       int64              `json:"expire" bson:"expire"`
	Items        []Product          `json:"items" bson:"items"`
	BusinessUnit string             `json:"businessUnit" bson:"businessUnit"`
	// UserData     User               `json:"userData" bson:"userData"`
	Agentsign string `json:"agentSign" bson:"agentSign"`
	User      string `json:"user" bson:"user,omitempty"`
}

// CartUpdate response structure after item update in cart
type CartUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Results       Cart  `json:"results"`
	NumberOfItems int64 `json:"numberOfItems"`
}

// RequestHeaders request headers
type RequestHeaders struct {
	UserToken    string `header:"gtw-sec-user-token" binding:"required"`
	BusinessUnit string `header:"gtw-business-unit" binding:"required"`
}

// ValidateExpiration validate expiry time
func (ct *Cart) ValidateExpiration() error {

	if ct.Expire <= 0 {
		return errors.New("invalid expiration time")
	}
	return nil
}

// SetCartExpiration set expiry time on cart
func (ct *Cart) SetCartExpiration() {
	ct.Expire = time.Now().UTC().Add(expirationTime * time.Hour).Unix()
}

// IsExpired checks cart expiration time
func (ct *Cart) IsExpired() bool {
	return time.Unix(ct.Expire, 0).Before(time.Now().UTC())
}

// SetUserData adds user data
/*func (ct *Cart) SetUserData(credential User) {
	ct.UserData = credential
}*/

// SetBusinessUnit adds business unit
func (ct *Cart) SetBusinessUnit(bu string) {
	ct.BusinessUnit = bu
}

// SetAgentSign sets user and agent sign
func (ct *Cart) SetAgentSign(credential User) {
	ct.Agentsign = credential.Credential.Agentsign
	ct.User = credential.Credential.User

}
