package dbconnection

// https://dev.to/eduardohitek/mongodb-golang-driver-tutorial-49e5
// https://github.com/eduardohitek/mongodb-go-example/blob/master/main.go
// https://github.com/tfogo/mongodb-go-tutorial/blob/master/main.go
// https://godoc.org/go.mongodb.org/mongo-driver/mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dburl = "mongodb://localhost:27017"
var database = "sunitestdb"
var collectionName = "products"

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
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		conn.l.Println("DB GetClient Connect :: ", err)
		log.Fatal(err)
	}
	return client
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
func InsertRecord(client *mongo.Client, product Product) interface{} {

	collection := client.Database(database).Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatalln("Error on inserting new product", err)
	}
	return insertResult.InsertedID
}

// InsertInterface is Temp Func for later implementation
func (conn *Connections) InsertInterface(client *mongo.Client) {
	// c := GetClient()
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
		conn.l.Println("Couldn't connect to the database", err)
	} else {
		conn.l.Println("Connected!")
	}

	product := Product{Fileid: 13, Position: 11, Filename: "Stephen Strange File", Description: "Doctor Strange", Filedate: "2021-01-21", Source: "From Hospital"}
	insertedID := InsertRecord(client, product)
	conn.l.Println("Record Inserted :: ", insertedID)
}
