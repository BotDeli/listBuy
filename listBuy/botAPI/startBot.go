package botAPI

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
	"listBuy/logger"
)

func StartBotAPI(telegramToken string, collection *mongo.Collection) {
	bot, update := GetBotAndUpdates(telegramToken)
	logger.Logger.Info("Successful start botAPI!")
	startUpdatesChannel(bot, update, collection)
}

func startUpdatesChannel(bot *tgbotapi.BotAPI, updatesChannel *tgbotapi.UpdatesChannel, collection *mongo.Collection) {
	for update := range *updatesChannel {
		MessageHandler(bot, &update, collection)
	}
}
