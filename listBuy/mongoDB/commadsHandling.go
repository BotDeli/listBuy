package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"listBuy/logger"
)

type UserInfo struct {
	UserId  int64    `bson:"userId" json:"userId"`
	ListBuy []string `bson:"listBuy" json:"listBuy"`
	FSM     string   `bson:"fSM" json:"fSM"`
}

func RegistrationUser(userID int64, collection *mongo.Collection) {
	userInfo := UserInfo{
		UserId:  userID,
		ListBuy: []string{},
		FSM:     "main",
	}
	_, err := collection.InsertOne(ctx, userInfo)
	logger.CheckWarning(err)
}

func GetFSM(userID int64, collection *mongo.Collection) string {
	userInfo := getUserInfo(userID, collection)
	return userInfo.FSM
}

func GetListBuy(userID int64, collection *mongo.Collection) []string {
	userInfo := getUserInfo(userID, collection)
	return userInfo.ListBuy
}

func getUserInfo(userID int64, collection *mongo.Collection) UserInfo {
	var user UserInfo
	cursor := findUser(userID, collection)
	if cursor.Next(ctx) {
		err := cursor.Decode(&user)
		logger.CheckWarning(err)
	}
	return user
}

func CheckUserInDB(userID int64, collection *mongo.Collection) bool {
	cursor := findUser(userID, collection)
	return cursor.Next(ctx)
}

func findUser(userID int64, collection *mongo.Collection) *mongo.Cursor {
	filter := getFilterUserID(userID)
	cursor, err := collection.Find(ctx, filter)
	logger.CheckWarning(err)
	return cursor
}

func ChangeFSM(userID int64, collection *mongo.Collection, newFSM string) {
	updateDataUser(userID, collection, "fSM", newFSM)
}

func AddItemInListBuy(userID int64, collection *mongo.Collection, item string) {
	listBuy := GetListBuy(userID, collection)
	listBuy = append(listBuy, item)
	updateDataUser(userID, collection, "listBuy", listBuy)
}

func RemoveItemInListBuy(userID int64, collection *mongo.Collection, itemID int) bool {
	listBuy := GetListBuy(userID, collection)
	if itemID > 0 && itemID <= len(listBuy) {
		listBuy = append(listBuy[:itemID-1], listBuy[itemID:]...)
		updateDataUser(userID, collection, "listBuy", listBuy)
		return true
	}
	return false
}

func updateDataUser(userID int64, collection *mongo.Collection, param string, value any) {
	filter := getFilterUserID(userID)
	update := getUpdate(param, value)
	_, err := collection.UpdateOne(ctx, filter, update)
	logger.CheckWarning(err)
}

func getFilterUserID(userId int64) bson.M {
	filter := bson.M{"userId": userId}
	return filter
}

func getUpdate(param string, value any) bson.M {
	update := bson.M{"$set": bson.M{param: value}}
	return update
}
