package repository

import (
    "context"
    "errors"
    "uas-backend-go/app/model"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type AchievementMongoRepository struct {
    Col *mongo.Collection
}

func NewAchievementMongoRepository(c *mongo.Collection) *AchievementMongoRepository {
    return &AchievementMongoRepository{Col: c}
}

//
// CREATE ACHIEVEMENT (Mahasiswa)
//
func (r *AchievementMongoRepository) Create(ctx context.Context, a model.Achievement) (primitive.ObjectID, error) {
    a.ID = primitive.NewObjectID()

    _, err := r.Col.InsertOne(ctx, a)
    return a.ID, err
}

//
// UPDATE ACHIEVEMENT
//
func (r *AchievementMongoRepository) Update(ctx context.Context, id primitive.ObjectID, a model.Achievement) error {
    upd := bson.M{
        "$set": a,
    }

    res, err := r.Col.UpdateByID(ctx, id, upd)
    if err != nil {
        return err
    }
    if res.MatchedCount == 0 {
        return errors.New("achievement not found")
    }
    return nil
}

//
// GET DETAIL
//
func (r *AchievementMongoRepository) GetByID(ctx context.Context, id primitive.ObjectID) (model.Achievement, error) {
    var a model.Achievement
    err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&a)
    return a, err
}

//
// SOFT DELETE
//
func (r *AchievementMongoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
    res, err := r.Col.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err
    }
    if res.DeletedCount == 0 {
        return errors.New("achievement not found")
    }
    return nil
}
