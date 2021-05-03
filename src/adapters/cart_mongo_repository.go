package adapters

import (
	"context"

	"github.com/mohsanabbas/ms-cart/src/domain/cart"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartMongoRepository struct {
	mongoClient *mongo.Client
}

func NewCartMongoRepository(
	mongoClient *mongo.Client,
) CartMongoRepository {
	return CartMongoRepository{
		mongoClient: mongoClient,
	}
}


func (r CartMongoRepository) cartDb() *mongo.Database {
	return r.mongoClient.Database("shopping-cart")
}
func (r CartMongoRepository) cartCollection() *mongo.Collection {
	return r.cartDb().Collection("cartte")
}



func (r CartMongoRepository) CreatCart(ctx context.Context, ct *cart.Cart) (error) {
// ct.Validate()
for k := range ct.Items {
	ct.Items[k].GenerateItemID()
}
	collection := r.cartCollection()
	_, err := collection.InsertOne(ctx, ct)
	if err != nil {
		return  errors.Wrap(err, "unable to get actual docs")
	}
	// id := res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}


//GetCart
func (r CartMongoRepository) GetCart(
	ctx context.Context,
	id string,
) (*cart.Cart, error) {

	cart:=cart.Cart{}

	cartID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to convert string into objectID")
	}
	collection := r.cartCollection()
	filter := bson.M{"_id": cartID}
	if err := collection.FindOne(ctx, filter).Decode(&cart); err != nil {
		return nil, errors.Wrap(err, "unable to get actual docs from mongo")
	}

return &cart, nil

}


func (r CartMongoRepository) AddProduct(
	ctx context.Context,
	id string,
	update *cart.Product,
	updateFn func(ctx context.Context, ct *cart.Cart) (*cart.Cart, error),
	) error {

		result := cart.CartUpdate{
			ModifiedCount: 0,
		}
	collection := r.cartCollection()
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "EROR ID-unable to get actual docs")
	}
	filter := bson.M{"$push": bson.M{"items": update}}
	res, err := collection.UpdateOne(ctx, bson.M{"_id": _id}, filter);
	if err != nil {
		return errors.Wrap(err, "updatation failed")
	}
	updated, err := updateFn(ctx, &cart.Cart{})
		if err != nil {
			return err
		}

	result.ModifiedCount=res.ModifiedCount
	result.Results =  *updated

	return  nil
}
