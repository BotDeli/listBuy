package main

import (
	"listBuy/botAPI"
	"listBuy/logger"
	"listBuy/mongoDB"
)

const telegramToken = "6142563783:AAFk9syFxKyCMFyPvDuW5ci4mK_FGv81i-0"
const uri = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.9.1"

func main() {
	logger.StartLogger()
	client, collection := mongoDB.ConnectDB(uri)
	defer mongoDB.DisconnectDB(client)
	botAPI.StartBotAPI(telegramToken, collection)
}
