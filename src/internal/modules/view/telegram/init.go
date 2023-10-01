package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"hospital/internal/modules/config"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	disease_dto "hospital/internal/modules/domain/disease/dto"
	patient_dto "hospital/internal/modules/domain/patient/dto"
	room_dto "hospital/internal/modules/domain/room/dto"
	"hospital/internal/modules/view/telegram/controllers"
	"strconv"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Зарегестрироваться"),
		tgbotapi.NewKeyboardButton("Помощь"),
		tgbotapi.NewKeyboardButton("Просмотреть данные о себе"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Удалить пациента"),
		tgbotapi.NewKeyboardButton("Добавить пациента"),
		tgbotapi.NewKeyboardButton("Посмотреть своих пациентов"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Найти палату по номеру"),
		tgbotapi.NewKeyboardButton("Вывести все палаты"),
		tgbotapi.NewKeyboardButton("Добавить палату"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Добавить заболевание"),
	),
)

func singUp(chatId int64, user *UsersMessage) string {
	var msg string
	addNextMessages("Введите свою фамилию", user, chatId)
	addNextMessages("Введите свою специальность", user, chatId)
	addNextMessages("Введите свою роль", user, chatId)
	msg, user.NextMessages = printNewMessage(user.NextMessages)
	return msg
}

func endSingUp(user *UsersMessage, chatId int64, controller *controllers.Controller) string {
	var reply string

	newDoctor := &auth_dto.NewDoctor{
		TokenId:    strconv.FormatInt(chatId, 10),
		Surname:    user.UserMessages[0],
		Speciality: user.UserMessages[1],
		Role:       user.UserMessages[2],
	}
	_, err := controller.SingUp(context.Background(), newDoctor)
	if err != nil {
		reply = "Уже зарегистрированы"
		return reply
	}
	reply = "Зарегистрирован"

	return reply
}

func endAddPatient(user *UsersMessage, controller *controllers.Controller) string {
	var reply string
	height, _ := strconv.Atoi(user.UserMessages[3])
	weight, _ := strconv.ParseFloat(user.UserMessages[4], 64)
	roomNumber, _ := strconv.Atoi(user.UserMessages[5])
	degreeOfDanger, _ := strconv.Atoi(user.UserMessages[6])

	newPatient := &patient_dto.CreatePatient{
		Surname:        user.UserMessages[0],
		Name:           user.UserMessages[1],
		Patronymic:     user.UserMessages[2],
		Height:         height,
		Weight:         weight,
		RoomNumber:     roomNumber,
		DegreeOfDanger: degreeOfDanger,
	}
	_, err := controller.AddPatient(context.Background(), newPatient)
	if err != nil {
		reply = "Ошбика добавления"
		return reply
	}
	reply = "Добавлен"

	return reply
}

func endAddRoom(user *UsersMessage, controller *controllers.Controller) string {
	var reply string

	num, _ := strconv.Atoi(user.UserMessages[0])
	floor, _ := strconv.Atoi(user.UserMessages[1])
	numberBeds, _ := strconv.Atoi(user.UserMessages[2])

	newRoom := &room_dto.CreateRoom{
		Num:            num,
		Floor:          floor,
		NumberBeds:     numberBeds,
		TypeRoom:       user.UserMessages[3],
		NumberPatients: 0,
	}
	_, err := controller.AddRoom(context.Background(), newRoom)
	if err != nil {
		reply = "Ошбика добавления комнаты"
		return reply
	}
	reply = "Комната добалена!"

	return reply
}

func endAddDisease(user *UsersMessage, controller *controllers.Controller) string {
	var reply string

	degreeOfDanger, _ := strconv.Atoi(user.UserMessages[1])

	newDisease := &disease_dto.CreateDisease{
		Name:           user.UserMessages[0],
		DegreeOfDanger: degreeOfDanger,
		Threat:         user.UserMessages[2],
	}
	_, err := controller.AddDisease(context.Background(), newDisease)
	if err != nil {
		reply = "Ошбика добавления заболевания"
		return reply
	}
	reply = "Заболевание добалено! еее"

	return reply
}

func handleUsers(
	Users []UsersMessage, chatId int64,
	userMessage string,
	controller *controllers.Controller) (string, []UsersMessage) {

	var msg string
	for i, u := range Users {
		if u.ChatId == chatId {
			Users[i].UserMessages = append(Users[i].UserMessages, userMessage)
			if len(u.NextMessages) > 0 {
				msg, Users[i].NextMessages = printNewMessage(u.NextMessages)
			} else {
				switch u.Command {
				case "Зарегестрироваться":
					msg = endSingUp(&Users[i], chatId, controller)
				case "Добавить пациента":
					msg = endAddPatient(&Users[i], controller)
				case "Добавить палату":
					msg = endAddRoom(&Users[i], controller)
				case "Добавить заболевание":
					msg = endAddDisease(&Users[i], controller)

				}
				Users = Users[:i+copy(Users[i:], Users[i+1:])]

			}
		}
	}
	return msg, Users
}

func getInfoAboutDoctor(id int64, controller *controllers.Controller) string {
	token := strconv.FormatInt(id, 10)
	doctor, err := controller.DoctorToken(context.Background(), token)
	if err != nil {
		msg := "Вы не зарегистрированы"
		return msg
	}
	msg := fmt.Sprintf("Фамилия: %s \nСпециальность: %s \nРоль: %s \n",
		doctor.Surname, doctor.Speciality, doctor.Role)
	return msg
}

func addPatient(chatId int64, user *UsersMessage) string {
	var msg string
	addNextMessages("Введите фамилию пациента", user, chatId)
	addNextMessages("Введите имя пациента", user, chatId)
	addNextMessages("Введите отчество пациента", user, chatId)
	addNextMessages("Введите Рост пациента", user, chatId)
	addNextMessages("Введите Вес пациента", user, chatId)
	addNextMessages("Введите номер комнаты пациента", user, chatId)
	addNextMessages("Введите степень опасности пациента", user, chatId)
	msg, user.NextMessages = printNewMessage(user.NextMessages)
	return msg
}

func addRoom(chatId int64, user *UsersMessage) string {
	var msg string
	addNextMessages("Введите Номер палаты", user, chatId)
	addNextMessages("Введите Этаж палаты", user, chatId)
	addNextMessages("Введите Количетво кроватей палаты", user, chatId)
	addNextMessages("Введите Тип палаты", user, chatId)
	msg, user.NextMessages = printNewMessage(user.NextMessages)
	return msg
}

func addDisease(chatId int64, user *UsersMessage) string {
	var msg string
	addNextMessages("Введите заболевание", user, chatId)
	addNextMessages("Введите степень опасности заболевания", user, chatId)
	addNextMessages("Введите способ лечения", user, chatId)
	msg, user.NextMessages = printNewMessage(user.NextMessages)
	return msg
}

func getInfoAboutPatients(id int64, controller *controllers.Controller) string {
	var msg string = ""
	patients, err := controller.GetAllPatients(context.Background())
	if err != nil {
		msg := "Ошбика запроса"
		return msg
	}

	for i := range patients {
		msg += fmt.Sprintf("ID %d \nФамилия: %s \nИмя: %s \nОтчество: %s \n",
			patients[i].Id, patients[i].Surname, patients[i].Name, patients[i].Patronymic)
	}
	return msg
}

func printAllRooms(id int64, controller *controllers.Controller) string {
	var msg string = ""
	rooms, err := controller.GetAllRooms(context.Background())
	if err != nil {
		msg := "Ошбика запроса"
		return msg
	}

	for i := range rooms {
		msg += fmt.Sprintf("ID %d \nНомер: %d \nЭтаж: %d \nТип: %s \n",
			rooms[i].Id, rooms[i].Num, rooms[i].Floor, rooms[i].TypeRoom)
	}
	return msg
}

func handleBot(
	controller *controllers.Controller,
	updates tgbotapi.UpdatesChannel,
	bot *tgbotapi.BotAPI,
	logger *zap.Logger) {

	var Users []UsersMessage

	for update := range updates {
		if update.Message != nil {

			ChatId := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(ChatId, update.Message.Text)

			logger.Info(msg.Text)

			if len(Users) > 0 {
				msg.Text, Users = handleUsers(Users, ChatId, update.Message.Text, controller)
			} else {
				switch update.Message.Text {
				case "Помощь":
					msg.Text = "Напишите: Зарегестрироваться или open"
				case "Зарегестрироваться":
					Users = append(Users, UsersMessage{ChatId: ChatId, Command: update.Message.Text})
					msg.Text = singUp(ChatId, &Users[len(Users)-1])
				case "Просмотреть данные о себе":
					msg.Text = getInfoAboutDoctor(ChatId, controller)
				case "Добавить пациента":
					Users = append(Users, UsersMessage{ChatId: ChatId, Command: update.Message.Text})
					msg.Text = addPatient(ChatId, &Users[len(Users)-1])
				case "Посмотреть своих пациентов":
					msg.Text = getInfoAboutPatients(ChatId, controller)
				case "Вывести все палаты":
					msg.Text = printAllRooms(ChatId, controller)
				case "Добавить палату":
					Users = append(Users, UsersMessage{ChatId: ChatId, Command: update.Message.Text})
					msg.Text = addRoom(ChatId, &Users[len(Users)-1])
				case "Добавить заболевание":
					Users = append(Users, UsersMessage{ChatId: ChatId, Command: update.Message.Text})
					msg.Text = addDisease(ChatId, &Users[len(Users)-1])
				case "open":
					msg.ReplyMarkup = numericKeyboard
				case "close":
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				default:
					msg.Text = "Команда не найдена, напишите Помощь"
					logger.Warn("Неверная пользовательская комманда")
				}
			}

			if _, err := bot.Send(msg); err != nil {
				logger.Error("Ошибка запроса")
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

func startBot(controller *controllers.Controller, cfg config.Config, logger *zap.Logger) {
	dotenv := cfg.TelegramToken

	bot, err := tgbotapi.NewBotAPI(dotenv)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	handleBot(controller, updates, bot, logger)
}
