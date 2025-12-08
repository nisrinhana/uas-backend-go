package database

import (
    "context"
    "fmt"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Client
var MongoDB *mongo.Database // <- ini penting

func InitMongo() error {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        return fmt.Errorf("MONGO_URI is empty")
    }

    dbName := os.Getenv("MONGO_DB")
    if dbName == "" {
        return fmt.Errorf("MONGO_DB is empty")
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := client.Connect(ctx); err != nil {
        return err
    }

    Mongo = client
    MongoDB = client.Database(dbName) // <- pilih database dari .env

    fmt.Println("Connected to MongoDB:", dbName)
    return nil
}
