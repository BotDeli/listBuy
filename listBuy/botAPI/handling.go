package botAPI

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
	"listBuy/logger"
	"listBuy/mongoDB"
	"strconv"
	"strings"
)

func MessageHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update, collection *mongo.Collection) {
	userID := GetUserID(update)
	FSM := mongoDB.GetFSM(userID, collection)
	if update.Message != nil {
		if update.Message.IsCommand() {
			CommandHandler(bot, update, collection)
		} else {
			switch FSM {
			case "add":
				addItem(bot, update, userID, collection)
			case "remove":
				removeItem(bot, update, userID, collection)
			}
		}
	} else if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case "addItem":
			mongoDB.ChangeFSM(userID, collection, "add")
			SendMessage(bot, update, "Введите название предмета, чтобы добавить его в список:", "")
		case "removeItem":
			mongoDB.ChangeFSM(userID, collection, "remove")
			SendMessage(bot, update, "Введите номер предмета, чтобы его удалить:", "")
		case "updateList":
			updateList(bot, update, userID, collection)
		default:
			mongoDB.ChangeFSM(userID, collection, "main")
		}
	}
}

func addItem(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userID int64, collection *mongo.Collection) {
	item := strings.TrimSpace(update.Message.Text)
	if item != "" {
		mongoDB.AddItemInListBuy(userID, collection, item)
		SendMessage(bot, update, "Успешно добавлено в список!", "")
		mongoDB.ChangeFSM(userID, collection, "main")
	}
}

func removeItem(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userID int64, collection *mongo.Collection) {
	message := strings.TrimSpace(update.Message.Text)
	itemID, err := strconv.Atoi(message)
	if err == nil {
		if mongoDB.RemoveItemInListBuy(userID, collection, itemID) {
			SendMessage(bot, update, "Успешно удалено из списка!", "")
		} else {
			SendMessage(bot, update, "Некорректный номер предмета", "")
		}
		mongoDB.ChangeFSM(userID, collection, "main")
	} else {
		logger.CheckWarning(err)
		SendMessage(bot, update, "Некорректный ввод", "")
	}
}

func updateList(bot *tgbotapi.BotAPI, update *tgbotapi.Update, userID int64, collection *mongo.Collection) {
	strListBuy := getStrListBuy(userID, collection)
	messageID := update.CallbackQuery.Message.MessageID
	editConfig := tgbotapi.NewEditMessageText(userID, messageID, strListBuy)
	keyboard := getMainMenuInlineKeyboard()
	editConfig.ReplyMarkup = keyboard
	_, err := bot.Send(editConfig)
	logger.CheckWarning(err)
}
