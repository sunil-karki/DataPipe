package dbconnection

// https://dev.to/eduardohitek/mongodb-golang-driver-tutorial-49e5
// https://github.com/eduardohitek/mongodb-go-example/blob/master/main.go
// https://github.com/tfogo/mongodb-go-tutorial/blob/master/main.go
// https://godoc.org/go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB Configs
var dburl = "mongodb://localhost:27017"
var database = "sunitestdb"
var collectionName = "products"

// This clientcon will be used to perform ping and crud operations like insert, update.......etc.
var clientcon *mongo.Client

// Connections is a MongoDB Connection handler
type Connections struct {
	l *log.Logger
}

// NewConnection creates a connection with the given logger
func NewConnection(l *log.Logger) *Connections {
	return &Connections{l}
}

// Connect creates and checks connection
func (conn *Connections) Connect() {
	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	// uri := "mongodb://localhost:27017/sunitestdb"

	conn.l.Println("Trying to connect to Mongod Location: ", dburl+"/"+database)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl+"/"+database))
	if err != nil {
		conn.l.Println("DB Connect :: ", err)
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			conn.l.Println("DB Disconnect :: ", err)
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		conn.l.Println("DB Ping :: ", err)
		panic(err)
	}

	conn.l.Println("Successfully connected and pinged...")
}

//GetClient returns a MongoDB Client
func (conn *Connections) GetClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI(dburl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		conn.l.Println("DB GetClient :: ", err)
		// log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		conn.l.Println("DB GetClient Connect :: ", err)
		// log.Fatal(err)
	}
	return client

}

// CreateConnection gets GetClient, checks connection and returns connection for CRUD Interface.
func (conn *Connections) CreateConnection() {

	// clientConn := conn.GetClient()
	// err := clientConn.Ping(context.Background(), readpref.Primary())
	// if err != nil {
	// 	log.Fatal("Couldn't connect to the database", err)
	// 	conn.l.Println("Couldn't connect to the database", err)
	// } else {
	// 	conn.l.Println("Connected!")
	// }

	clientcon = conn.GetClient()
	err := clientcon.Ping(context.Background(), readpref.Primary())
	if err != nil {
		// log.Fatal("Couldn't connect to the database", err)
		conn.l.Println("Couldn't connect to the database", err)
	} else {
		conn.l.Println("Connected :: Connection Created!")
	}

}

// Product defines the structure for an API product
type Product struct {
	Fileid      int    `json:"fileid" bson:"fileid"`
	Position    int    `json:"position" bson:"position"`
	Filename    string `json:"filename" bson:"filename"`
	Description string `json:"description" bson:"description"`
	Filedate    string `json:"filedate" bson:"filedate"`
	Source      string `json:"source" bson:"source"`
}

// InsertRecord insert a new record in the Collection
// Returns ObjectID of each Document/Record inserted.
func (conn *Connections) InsertRecord(product Product) interface{} {

	collection := clientcon.Database(database).Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		conn.l.Println("Error on inserting new product", err)
	}
	return insertResult.InsertedID
}

// InsertInterface is Temp Func for later implementation
func (conn *Connections) InsertInterface() {

	product := Product{Fileid: 14, Position: 12, Filename: "Doctor Strange File", Description: "Doctor Strange", Filedate: "2021-01-24", Source: "From PY"}
	// insertedID := InsertRecord(clientcon, product)
	insertedID := conn.InsertRecord(product)
	conn.l.Println("Record Inserted :: ", insertedID)

}

// UpdateRecord updates the value of a record
// Returns no. of record updated as Int .
func (conn *Connections) UpdateRecord(updatedData interface{}, where bson.M) int64 {

	collection := clientcon.Database(database).Collection(collectionName)
	setBy := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), where, setBy)
	if err != nil {
		conn.l.Println("Error on updating one Record", err)
	}
	return updatedResult.ModifiedCount
}

// UpdateInterface is Temp Func for later implementation
func (conn *Connections) UpdateInterface() {
	updatedCnt := conn.UpdateRecord(bson.M{"filedate": "2021-03-24"}, bson.M{"source": "From PY"})
	conn.l.Println("Record Updated count:", updatedCnt)
}

// DeleteRecord remove one existing Record
// Returns no. of Record/Document deleted.
func (conn *Connections) DeleteRecord(filter bson.M) int64 {

	collection := clientcon.Database(database).Collection(collectionName)
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		conn.l.Println("Error on deleting one Record", err)
	}
	return deleteResult.DeletedCount
}

// DeleteInterface is Temp Func for later implementation
func (conn *Connections) DeleteInterface() {
	deletedCnt := conn.DeleteRecord(bson.M{"filedate": "2021-03-24"})
	conn.l.Println("Record Removed count:", deletedCnt)
}

// ReturnAllRecords return all documents from the collection Products
func (conn *Connections) ReturnAllRecords(filter bson.M) []*Product {
	var products []*Product
	collection := clientcon.Database(database).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		conn.l.Println("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var product Product
		err = cur.Decode(&product)
		if err != nil {
			conn.l.Println("Error on Decoding the document", err)
		}
		products = append(products, &product)
	}
	return products
}

// ReturnRecordInterface is Temp Func for later implementation
func (conn *Connections) ReturnRecordInterface() {
	conn.l.Println("Returning All")
	products := conn.ReturnAllRecords(bson.M{})
	for _, product := range products {
		conn.l.Println(product.Filename, product.Description, product.Filedate)
	}

	conn.l.Println("Returning matching Filedate")
	products = conn.ReturnAllRecords(bson.M{"filedate": "2021-02-24"})
	for _, product := range products {
		conn.l.Println(product.Filename, product.Description, product.Filedate)
	}
}
