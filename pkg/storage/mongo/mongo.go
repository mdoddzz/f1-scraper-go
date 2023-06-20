package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Storage represents a storage object
type Storage struct {
	db                           *mongo.Database
	race                         *mongo.Collection
	race_results                 *mongo.Collection
	driver_standings_season      *mongo.Collection
	constructor_standings_season *mongo.Collection
	qualifying                   *mongo.Collection
	pit_stops                    *mongo.Collection
	starting_grid                *mongo.Collection
}

// MongoDB order
type Order struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

// NewDB initializes the database connection
func NewDB(uri string, dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// test connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbName)
}

// NewStorage returns a new instance of Storage
func NewStorage(db *mongo.Database) *Storage {

	s := Storage{
		db:                           db,
		race:                         db.Collection("race"),
		race_results:                 db.Collection("race_results"),
		driver_standings_season:      db.Collection("driver_standings_season"),
		constructor_standings_season: db.Collection("constructor_standings_season"),
		qualifying:                   db.Collection("qualifying"),
		pit_stops:                    db.Collection("pit_stops"),
		starting_grid:                db.Collection("starting_grid"),
	}

	return &s
}

/*
func findOptions(offset int, limit int, order []Order) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit)).
		SetSkip(int64(offset))

	sort := bson.D{}
	for _, o := range order {
		dir := 1
		if o.Dir == "DESC" || o.Dir == "desc" {
			dir = -1
		}
		sort = append(sort, primitive.E{Key: o.Name, Value: dir})
	}

	findOptions.SetSort(sort)

	return findOptions
}
*/
