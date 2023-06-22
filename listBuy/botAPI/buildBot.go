package botAPI

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"listBuy/logger"
)

func GetBotAndUpdates(telegramToken string) (*tgbotapi.BotAPI, *tgbotapi.UpdatesChannel) {
	bot := createBot(telegramToken)
	updateConfig := getUpdateConfig()
	updatesChannel, err := bot.GetUpdatesChan(updateConfig)
	logger.CheckError(err)
	return bot, &updatesChannel
}

func createBot(telegramToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	logger.CheckError(err)
	return bot
}

func getUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	return updateConfig
}
