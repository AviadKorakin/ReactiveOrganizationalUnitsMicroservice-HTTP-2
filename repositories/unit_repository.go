package repositories

import (
	"context"
	"time"

	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UnitRepository struct {
	coll *mongo.Collection
}

func NewUnitRepository(client *mongo.Client) *UnitRepository {
	return &UnitRepository{coll: client.Database("orgdb").Collection("units")}
}

func (r *UnitRepository) Create(ctx context.Context, u *models.Unit) error {
	u.CreationDate = time.Now().Format("02-01-2006")
	if u.UnitID == "" {
		u.UnitID = bson.NewObjectID().Hex()
	}
	_, err := r.coll.InsertOne(ctx, u)
	return err
}

func (r *UnitRepository) GetByID(ctx context.Context, id string) (*models.Unit, error) {
	var u models.Unit
	err := r.coll.FindOne(ctx, bson.M{"unitId": id}).Decode(&u)
	return &u, err
}

func (r *UnitRepository) List(ctx context.Context, page, size int64) ([]models.Unit, error) {
	opts := options.Find().
		SetSkip((page - 1) * size).
		SetLimit(size).
		SetSort(bson.D{{Key: "unitId", Value: 1}})
	cursor, err := r.coll.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var units []models.Unit
	if err := cursor.All(ctx, &units); err != nil {
		return nil, err
	}
	return units, nil
}

func (r *UnitRepository) Update(ctx context.Context, id string, upd bson.M) error {
	_, err := r.coll.UpdateOne(ctx, bson.M{"unitId": id}, bson.M{"$set": upd})
	return err
}

func (r *UnitRepository) DeleteAll(ctx context.Context) error {
	_, err := r.coll.DeleteMany(ctx, bson.D{})
	return err
}
