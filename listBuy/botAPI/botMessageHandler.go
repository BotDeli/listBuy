package botAPI

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"listBuy/logger"
)

func SendMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update, text string, addition string) {
	userID := GetUserID(update)
	message := tgbotapi.NewMessage(userID, text)
	if addition != "" {
		switchKeyboard(&message, addition)
	}
	_, err := bot.Send(message)
	logger.CheckWarning(err)
}

func RemoveMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	userID := GetUserID(update)
	messageID := GetMessageID(update)
	deleteMessageConfig := tgbotapi.NewDeleteMessage(userID, messageID)
	_, err := bot.Send(deleteMessageConfig)
	logger.CheckWarning(err)
}

func GetUserID(update *tgbotapi.Update) int64 {
	if update.Message != nil {
		return update.Message.Chat.ID
	}
	return update.CallbackQuery.Message.Chat.ID
}

func GetMessageID(update *tgbotapi.Update) int {
	if update.Message != nil {
		return update.Message.MessageID
	}
	return update.CallbackQuery.Message.MessageID
}

func switchKeyboard(message *tgbotapi.MessageConfig, addition string) {
	switch addition {
	case "main":
		menuKeyboard := getMainMenuInlineKeyboard()
		(*message).ReplyMarkup = menuKeyboard
	}
}
