package dbconnection

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
	uri := "mongodb://localhost:27017/sunitestdb"

	conn.l.Println("Trying to connect to Mongod Location: ", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
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
