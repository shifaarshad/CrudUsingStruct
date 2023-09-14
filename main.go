package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string
	Email string
}

func main() {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb+srv://shifahajiarshad:hajiarshad99@cluster0.bfkb4kx.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB Atlas:", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	collection := client.Database("practice").Collection("crud")

	// Create a new user
	user := User{Name: "jamal", Email: "john@example.com"}
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println("Failed to insert user:", err)
		return
	}

	// Create a new user
	newuser := User{Name: "ahsan", Email: "john@example.com"}
	_, err = collection.InsertOne(context.Background(), newuser)
	if err != nil {
		fmt.Println("Failed to insert user:", err)
		return
	}
	// Retrieve a user by email
	var result User
	err = collection.FindOne(context.Background(), bson.M{"email": "john@example.com"}).Decode(&result)
	if err != nil {
		fmt.Println("Failed to find user:", err)
		return
	}
	fmt.Println("Retrieved user:", result)

	// Update a user's name
	update := bson.M{"$set": bson.M{"name": "John Doe"}}
	_, err = collection.UpdateOne(context.Background(), bson.M{"email": "john@example.com"}, update)
	if err != nil {
		fmt.Println("Failed to update user:", err)
		return
	}

}
