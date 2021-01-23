package dbconnection

import (
	"context"
	"fmt"
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

func (conn *Connections) connection() {
	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	uri := "mongodb://localhost:27017/sunitestdb"

	conn.l.Println("Trying to connect to Mongod Location: ", uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	conn.l.Println("Successfully connected and pinged...")
}
