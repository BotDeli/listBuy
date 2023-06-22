package botAPI

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func getMainMenuInlineKeyboard() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Добавить в список покупок",
				"addItem",
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Удалить из списка покупок",
				"removeItem",
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Обновить список покупок",
				"updateList",
			),
		),
	)
	return &keyboard
}
