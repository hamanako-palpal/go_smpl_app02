package util

import (
	"context"
	"fmt"
	"log"

	"github.com/hamanako-palpal/go_smpl_app02/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUser = "haruhito0823"
const cluster = "smplcluster-xfhz0.gcp.mongodb.net"

// InsertOne 一件追加
func InsertOne(us *entity.User) error {

	collection := dbConnecton()
	_, err := collection.InsertOne(context.Background(), us)

	return err
}

//Select ooo
func Select() {

}

// dbConnecton DBアクセッサ
func dbConnecton() *mongo.Collection {

	dburl := fmt.Sprintf("mongodb+srv://dbUser:%s@%s/admin?retryWrites=true&w=majority", dbUser, cluster)

	client, err := mongo.NewClient(options.Client().ApplyURI(dburl))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		return nil
	}
	defer client.Disconnect(context.Background())
	return client.Database("smpl").Collection("test_curry")
}
