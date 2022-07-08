package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func CheckError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func insertData(coll *mongo.Collection) {

	mahesh := Person{"Mahesh", 26}
	vinod := Person{"Vinod", 59}
	aruna := Person{"Aruna", 48}

	//Insert data in DB
	//Insert one row
	_, eio := coll.InsertOne(context.TODO(), mahesh)
	CheckError(eio)
	//Insert Multipe row
	persons := []interface{}{vinod, aruna}
	_, eim := coll.InsertMany(context.TODO(), persons)
	CheckError(eim)
}

func updateData(coll *mongo.Collection) {

	filter := bson.D{primitive.E{Key: "name", Value: "Mahesh"}}

	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "age", Value: "25"},
		}},
	}
	_, e := coll.UpdateOne(context.TODO(), filter, update)
	CheckError(e)

}

func findData(coll *mongo.Collection) {
	var res Person
	filter := bson.D{primitive.E{Key: "name", Value: "Mahesh"}}

	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	CheckError(err)
	fmt.Println(res)
}

func deleteData(coll *mongo.Collection) {
	// deleteFilter := bson.D{primitive.E{Key: "name", Value: "Jane"},
	// 	primitive.E{Key: "name", Value: "Ben"}}

	_, e := coll.DeleteMany(context.TODO(), bson.D{{}})
	CheckError(e)
}

func main() {
	//set client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	//create client from options ie. connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	CheckError(err)

	//ping the Db
	er := client.Ping(context.TODO(), nil)
	CheckError(er)

	coll := client.Database("testDb").Collection("firstCollection")

	insertData(coll)
	//updateData(coll)
	//findData(coll)
	//deleteData(coll)

}
