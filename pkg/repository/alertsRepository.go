package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
	"walmartTest/pkg/models"
)

var collection = new(mongo.Collection)

const AlertsCollection = "Alerts"

type AlarmRepository struct{}

func InitMongoDatabase() {
	log.Println("Initiating MongoDb connection")
	ctx, _ := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://test:password@127.0.0.1:10001"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Connected to MongoDb")
	collection = client.Database("WalmartTest").Collection(AlertsCollection)
}

func (p *AlarmRepository) Insert(alert models.Alert) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	alert.CreatedTimeStamp = time.Now()
	result, err := collection.InsertOne(ctx, &alert)
	if err != nil {
		log.Fatal("Failed to Insert Document", err)
		return nil, err
	}
	log.Println("Inserted document: ")
	return result, err
}

func getTimeFromEpochTime(epochTime string) time.Time {
	timestamp, err := strconv.ParseInt(epochTime, 10, 64)
	if err != nil {
		panic(err)
	}
	t := time.Unix(timestamp, 0)
	return t
}
func (p *AlarmRepository) Find(getAlertsRequest models.GetAlertsRequest) ([]*models.Alert, error) {
	var results []*models.Alert

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{
		"serviceid": getAlertsRequest.ServiceId,
		"createdtimestamp": bson.M{
			"$gt": primitive.NewDateTimeFromTime(getTimeFromEpochTime(getAlertsRequest.StartTs)),
			"$lt": primitive.NewDateTimeFromTime(getTimeFromEpochTime(getAlertsRequest.EndTs)),
		},
	}
	log.Println("Find Filter: ")
	log.Println(filter)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	// Iterating through the cursor
	for cur.Next(ctx) {
		var elem *models.Alert
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, err
}
