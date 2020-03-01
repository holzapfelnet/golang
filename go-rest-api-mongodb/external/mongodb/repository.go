package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	repository "../../repository"
)

// MyMongoRepository contains the mongo database, the name of the collection and
// the connection status.
type MyMongoRepository struct {
	db         *mongo.Database
	Collection string
	Connected  bool
}

// NewMyMongoRepository creates a new MongoDB repository
func NewMyMongoRepository(host string, port int, username string, password string, database string, collection string) *MyMongoRepository {
	repo := MyMongoRepository{Collection: collection}

	credentials := ""
	if username != "" {
		credentials = fmt.Sprintf("%v:%v", username, password)
	}

	uri := fmt.Sprintf("mongodb://%v@%v:%v", credentials, host, port)
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

	repo.db = client.Database(database)
	repo.Connected = true
	return &repo
}

// AddItem adds a new ToDo item to the collection
func (r *MyMongoRepository) AddItem(data repository.ToDoItem) error {

	collection := r.db.Collection(r.Collection)

	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil

}

// GetAllItems returns all items stored within the collection
func (r *MyMongoRepository) GetAllItems() ([]*repository.ToDoItem, error) {

	collection := r.db.Collection(r.Collection)
	filter := bson.D{{}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var result []*repository.ToDoItem

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem repository.ToDoItem
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		result = append(result, &elem)
	}

	return result, nil

}
