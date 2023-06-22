package botAPI

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
	"listBuy/mongoDB"
)

func CommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update, collection *mongo.Collection) {
	userID := GetUserID(update)
	switch update.Message.Text {
	case "/start":
		if registrationUser(userID, collection) {
			SendMessage(bot, update, "Инструкция...", "")
		}
		strListBuy := getStrListBuy(userID, collection)
		SendMessage(bot, update, strListBuy, "main")
		mongoDB.ChangeFSM(userID, collection, "main")
	default:
		RemoveMessage(bot, update)
	}
}

func registrationUser(userID int64, collection *mongo.Collection) bool {
	userAlreadyRegistered := mongoDB.CheckUserInDB(userID, collection)
	if !userAlreadyRegistered {
		mongoDB.RegistrationUser(userID, collection)
		return true
	}
	return false
}

func getStrListBuy(userID int64, collection *mongo.Collection) string {
	listBuy := mongoDB.GetListBuy(userID, collection)
	if len(listBuy) > 0 {
		strListBuy := "Ваш список покупок:\n"
		for n, item := range listBuy {
			strListBuy += fmt.Sprintf("%d. %s\n", n+1, item)
		}
		return strListBuy
	}
	return "Ваш список покупок пуст..."
}
