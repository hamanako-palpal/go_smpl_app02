package util

import (
	"context"
	"fmt"
	"log"

	"github.com/hamanako-palpal/go_smpl_app02/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbCd = "mongodb+srv://%s:%s@%s/test?retryWrites=true&w=majority"
const dbUser = "1385.aso@gmail.com"
const pass = "haruhito0823"
const cluster = "smplcluster-xfhz0.gcp.mongodb.net"

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

	url := fmt.Sprintf(dbCd, dbUser, pass, cluster)
	log.Printf(url)

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(url),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())
	return client.Database("smpl").Collection("test_curry")
}
