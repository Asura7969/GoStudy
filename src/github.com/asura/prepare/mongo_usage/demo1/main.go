package main

// 启动mongodb ./mongod --dbpath /data1/mongodb-4.0/data/ --logpath /data1/mongodb-4.0/log/log.log --bind_ip 0.0.0.0 &
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}
type LogRecord struct {
	JobName   string    `bson:"jobName"`
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"timePoint"`
}

func main() {

	var (
		client *mongo.Client
		err    error
	)

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://47.99.61.133:27017")
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	record := &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Content:   "hello",
		Err:       "",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	collection := client.Database("my_db").Collection("my_collection")

	result, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
	}
	if result == nil {
		fmt.Println("nil")
	}

	fmt.Println(result.InsertedID)

}

type mgo struct {
	uri        string
	database   string
	collection string
}

func (m *mgo) Connect() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 在调用WithTimeout之后defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	if err != nil {
		log.Print(err)
	}
	collection := client.Database(m.database).Collection(m.collection)
	return collection
}
