package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/aditya2305/REST_API/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const collName = "watchlist"
const dbname = "Cluster0"

var collection *mongo.Collection

// connecting with MONGODB

func init() {
	// client_options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to MDB
	client, err := mongo.Connect(context.TODO(), clientOption) // contexting the connection for use in later
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbname).Collection(collName)

	fmt.Println("Collection is ready!")
}

// MONGODB helpers

// insert a record

func insertOneMovie(movie model.Netflix) {
	data, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a movie: ", data.InsertedID)
}

// update a record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}                       // creating the query for filter
	update := bson.M{"$set": bson.M{"watched": true}} // creating the query for updates

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie updated! ", res.ModifiedCount)
}

// delete one record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted! ", res.DeletedCount)
}

// delete multiple_records

func deleteManyMovies() int {
	filter := bson.D{{}} // it means delete everything

	res, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted! ", res.DeletedCount)
	return int(res.DeletedCount)
}

// getting all the data

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// controllers

func GetAllMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var movie model.Netflix

	_ = json.NewDecoder(req.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	params := mux.Vars(req) // getting all the params
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeleteAMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	params := mux.Vars(req)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeleteAllMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	count := deleteManyMovies()
	json.NewEncoder(w).Encode(count)
}
