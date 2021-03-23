package mg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMgoURI(user, password, address, database,mechanism,replicaset string) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?replicaSet=%s", user, password, address, database, replicaset)
	if mechanism != "" {
		uri += "?authMechanism=" + mechanism
	}
	return uri
}

func Client() {
	ctx := context.TODO()
	uri := GetMgoURI("stark","6aAMQ_YGy1hfIh2", "rs001-01.dev-mongo.bkjk.cn:18036,rs001-02.dev-mongo.bkjk.cn:18036,rs001-03.dev-mongo.bkjk.cn:18036","stark", "", "rs001")
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err!=nil {
		panic(err)
	}
	cur, err := c.Database("stark").Collection("riskm").Find(ctx, &bson.D{}, options.Find().SetLimit(10))
	if err!=nil {
		panic(err)
	}
	for cur.Next(ctx) {
		elem := map[string]interface{}{}
		err := cur.Decode(elem)
		if err!=nil {
			panic(err)
		}
		fmt.Println(elem)
	}
}