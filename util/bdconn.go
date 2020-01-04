package util

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hamanako-palpal/go_smpl_app02/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DbConn 一件追加
func DbConn(us *entity.User) error {

	collection := dbConnecton()
	_, err := collection.InsertOne(context.Background(), us)

	return err
}

//Select ooo
func Select() {

}

// dbConnecton DBアクセッサ
func dbConnecton() *mongo.Collection {

	cfg := GetConfig()

	url := fmt.Sprintf(
		cfg.Db.Address,
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.Cluster)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(url),
	)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	return client.Database("smpl").Collection("test_curry")
}
