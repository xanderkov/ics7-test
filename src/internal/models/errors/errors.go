package errors

type Const string

func (e Const) Error() string {
	return string(e)
}

const (
	ErrIdValidate = Const("неверно указан Id объекта")

	ErrInvalidToken = Const("некорректный токен")

	MsgBadRequest = "ошибка параметров запроса"
	ErrBadRequest = Const(MsgBadRequest)

	MsgJsonUnMarshal = "не удалось декодировать JSON"
	ErrJsonUnMarshal = Const(MsgJsonUnMarshal)
	MsgJsonMarshal   = "не удалось упаковать данные в JSON"
	ErrJsonMarshal   = Const(MsgJsonMarshal)

	ErrDatabaseRecordNotFound = Const("запись не найдена")
	ErrUniqueViolation        = Const("нарушение уникальности ключа")

	ErrAccessDenied = Const("недостаточно прав")
)
