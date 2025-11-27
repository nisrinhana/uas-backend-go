package repository

import (
    "context"
    "time"

    "uas-backend-go/app/model"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type AchievementMongoRepository struct {
    Collection *mongo.Collection
}

func NewAchievementMongoRepository(db *mongo.Database) *AchievementMongoRepository {
    return &AchievementMongoRepository{
        Collection: db.Collection("achievements"),
    }
}

func (r *AchievementMongoRepository) Create(ctx context.Context, ach *model.Achievement) (*mongo.InsertOneResult, error) {
    ach.ID = primitive.NewObjectID()
    ach.CreatedAt = time.Now().Unix()
    ach.UpdatedAt = time.Now().Unix()

    return r.Collection.InsertOne(ctx, ach)
}

func (r *AchievementMongoRepository) GetByID(ctx context.Context, id string) (model.Achievement, error) {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return model.Achievement{}, err
    }

    var achievement model.Achievement
    err = r.Collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&achievement)

    return achievement, err
}

func (r *AchievementMongoRepository) Update(ctx context.Context, id string, update bson.M) error {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    update["updated_at"] = time.Now().Unix()

    _, err = r.Collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": update})
    return err
}

func (r *AchievementMongoRepository) SoftDelete(ctx context.Context, id string) error {
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = r.Collection.UpdateOne(
        ctx,
        bson.M{"_id": oid},
        bson.M{"$set": bson.M{"deleted": true}},
    )

    return err
}

func (r *AchievementMongoRepository) List(ctx context.Context, filter bson.M) ([]model.Achievement, error) {
    cursor, err := r.Collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }

    var achievements []model.Achievement
    err = cursor.All(ctx, &achievements)

    return achievements, err
}
