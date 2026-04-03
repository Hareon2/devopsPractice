package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	mongoDB := MongoDB{}

	got := mongoDB.filter("test")

	want := bson.D{{
		Key: "skills.name",
		Value: bson.D{{
			Key:   "$regex",
			Value: "^test.*$",
		}, {
			Key:   "$options",
			Value: "i",
		}},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("filter() got = %v, want %v", got, want)
	}
}
