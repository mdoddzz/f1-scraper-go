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
	constructor_standings_season *mongo.Collection
	constructor_standings        *mongo.Collection
	driver_standings_season      *mongo.Collection
	driver_standings             *mongo.Collection
	fastest_laps_awards          *mongo.Collection
	fastest_laps                 *mongo.Collection
	pit_stops                    *mongo.Collection
	practices                    *mongo.Collection
	qualifying                   *mongo.Collection
	race_results                 *mongo.Collection
	race                         *mongo.Collection
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
		constructor_standings_season: db.Collection("constructor_standings_season"),
		constructor_standings:        db.Collection("constructor_standings"),
		driver_standings_season:      db.Collection("driver_standings_season"),
		driver_standings:             db.Collection("driver_standings"),
		fastest_laps_awards:          db.Collection("fastest_laps_awards"),
		fastest_laps:                 db.Collection("fastest_laps"),
		pit_stops:                    db.Collection("pit_stops"),
		practices:                    db.Collection("practices"),
		qualifying:                   db.Collection("qualifying"),
		race_results:                 db.Collection("race_results"),
		race:                         db.Collection("race"),
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
