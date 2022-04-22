package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client     *mongo.Client
	context    context.Context
	cancelFunc context.CancelFunc
}

// Close This method closes mongoDB connection and cancel context.
func (db *MongoDB) Close() {
	if db.cancelFunc != nil {
		defer db.cancelFunc()
	}
	defer func() {
		if err := db.client.Disconnect(db.context); err != nil {
			panic(err)
		}
	}()
}

// Connect This is a user defined method that returns
func (db *MongoDB) Connect(uri string) error {
	//db.context, db.cancelFunc = context.WithTimeout(context.Background(), 30*time.Second)
	db.context = context.Background()
	client, err := mongo.Connect(db.context, options.Client().ApplyURI(uri))
	db.client = client
	return err
}

// Ping This is a user defined method that accepts
func (db *MongoDB) Ping() error {
	if err := db.client.Ping(db.context, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// InsertOne is a user defined method, used to insert
func (db *MongoDB) InsertOne(dataBase string, col string, doc interface{}) (interface{}, error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(db.context, doc)
	return result.InsertedID, err
}

// ReplaceOne is a user defined method, used to insert
func (db *MongoDB) ReplaceOne(dataBase string, col string, query interface{}, doc interface{}) (interface{}, error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err := collection.ReplaceOne(db.context, query, doc)
	return result.UpsertedID, err
}

// UpdateOne is a user defined method, used to insert
func (db *MongoDB) UpdateOne(dataBase string, col string, query interface{}, doc interface{}) (interface{}, error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err := collection.UpdateOne(db.context, query, doc)
	return result.UpsertedID, err
}

// InsertMany is a user defined method, used to insert
func (db *MongoDB) InsertMany(dataBase, col string, docs []interface{}) ([]interface{}, error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err := collection.InsertMany(db.context, docs)
	return result.InsertedIDs, err
}

// DeleteOne is a user defined function that delete,
func (db *MongoDB) DeleteOne(dataBase, col string, query interface{}) (int64, error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err := collection.DeleteOne(db.context, query)
	return result.DeletedCount, err
}

// DeleteMany is a user defined function that delete,
func (db *MongoDB) DeleteMany(dataBase string, col string, query interface{}) (int64, error) {
	collection := db.client.Database(dataBase).Collection(col)
	result, err := collection.DeleteMany(db.context, query)
	return result.DeletedCount, err
}

// Query method returns a cursor and error.
func (db *MongoDB) Query(dataBase string, col string, query interface{}, field interface{}) ([]bson.M, error) {
	collection := db.client.Database(dataBase).Collection(col)
	cursor, err := collection.Find(db.context, query, options.Find().SetProjection(field))
	defer cursor.Close(context.TODO())
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results, err
}
