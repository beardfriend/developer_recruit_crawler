package infrastructure

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMongodbConnection(t *testing.T) {
	m := NewMongodb("mongodb://username:password1234@localhost:27017/?retryWrites=true&w=majority")

	err := m.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}
}
