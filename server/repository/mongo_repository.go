package repository

import (
	"context"

	"github.com/mohammadzaidhussain/bookshop/config"
	"github.com/mohammadzaidhussain/bookshop/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoRepository interface {
	Create(data interface{}, ctx mongo.SessionContext) (interface{}, error)
	FindOne(id string, ctx mongo.SessionContext) (interface{}, error)
	FindOneByKey(key string, value interface{}, ctx mongo.SessionContext) (interface{}, error)
	Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error)
	Delete(id string, ctx mongo.SessionContext) (interface{}, error)
	FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error)
	Aggregate(pipelines primitive.A, ctx mongo.SessionContext) ([]map[string]interface{}, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func getSessionContext(sessionContext mongo.SessionContext) mongo.SessionContext {
	// cont, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	cont := context.Background()
	if sessionContext == nil {
		return mongo.NewSessionContext(cont, mongo.SessionFromContext(cont))
	}
	return sessionContext
}

func (mr *MongoRepository) Create(data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	result, err := mr.collection.InsertOne(sessionContext, data)
	return result, err
}

func (mr *MongoRepository) FindOne(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := mr.collection.FindOne(sessionContext, bson.M{"_id": objId})
	var document map[string]interface{}
	if err := result.Decode(&document); err != nil {
		return nil, err
	}
	return document, nil
}

func (mr *MongoRepository) FindOneByKey(key string, value interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(value.(string))
	var result *mongo.SingleResult
	if err != nil {
		//fmt.Printf("%s is not an object id", key)
		result = mr.collection.FindOne(sessionContext, bson.M{key: value})
	} else {
		result = mr.collection.FindOne(sessionContext, bson.M{key: objId})
	}
	var document map[string]interface{}
	if err := result.Decode(&document); err != nil {
		return nil, err
	}
	return document, nil
}

func (mr *MongoRepository) Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, _ := primitive.ObjectIDFromHex(id)
	res, err := mr.collection.UpdateOne(sessionContext, bson.M{"_id": objId}, bson.M{"$set": data})
	return res, err
}

func (mr *MongoRepository) Delete(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, _ := primitive.ObjectIDFromHex(id)
	res, err := mr.collection.DeleteOne(sessionContext, bson.M{"_id": objId})
	return res, err
}

func (mr *MongoRepository) FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getSessionContext(ctx)
	var finalFilter = bson.M{}
	if filter == nil {
		filter = finalFilter
	} else {
		staffDto := filter.(dto.StaffFilter)
		if staffDto.BookshopId != nil {
			finalFilter["bookshop_id"] = staffDto.BookshopId
		}
	}
	cursor, err := mr.collection.Find(sessionContext, finalFilter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(sessionContext)

	var results []map[string]interface{}
	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		results = append(results, document)
	}
	return results, cursor.Err()

}

func (mr *MongoRepository) Aggregate(pipelines primitive.A, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getSessionContext(ctx)
	cursor, err := mr.collection.Aggregate(sessionContext, pipelines)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(sessionContext)

	var results []map[string]interface{}
	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		results = append(results, document)
	}
	return results, cursor.Err()
}

func GetMongoRepository(dbName string, collectionName string) IMongoRepository {
	collection := config.GetDatabaseCollection(&dbName, collectionName)
	return &MongoRepository{
		collection: collection,
	}
}
