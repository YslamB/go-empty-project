package model

var InternalServerError = ServerError{
	Message: ResultMessage{
		Tk: "Serwer ýalňyşlygy",
		Ru: "Ошибка сервера",
		En: "Internal server error",
	},
}

var InvalidInput = ServerError{
	Message: ResultMessage{
		Tk: "Nädogry maglumat",
		Ru: "Неверное тело запроса",
		En: "Invalid request body",
	},
}

var Forbidden = ServerError{
	Message: ResultMessage{
		Tk: "Yalnysh maglumat",
		Ru: "Неверное тело запроса",
		En: "Invalid request body",
	},
}

var Unauthorized = ServerError{
	Message: ResultMessage{
		Tk: "Hesap döredilmedik",
		Ru: "Аккаунт не создан",
		En: "Account not created",
	},
}

var ServiceUnavailableWait = ResultMessage{
	Tk: "Hayyş garaşyň",
	Ru: "пожалуйста, подождите",
	En: "please wait",
}

var UnauthorizedError = ServerError{
	Message: ResultMessage{
		Tk: "Hesap döredilmedik",
		Ru: "Аккаунт не создан",
		En: "Account not created",
	},
}

var Forbitten = ServerError{
	Message: ResultMessage{
		Tk: "Abunaňyz ýok, yada çäkden geçdiňiz",
		Ru: "У вас нет подписки",
		En: "You don't have a subscription",
	},
}

var Conflict = ServerError{
	Message: ResultMessage{
		Tk: "MAglumat eýýäm bar",
		Ru: "уже существует",
		En: "already exists",
	},
}

var NotFound = ServerError{
	Message: ResultMessage{
		Tk: "Maglumat Tapylmady",
		Ru: "Ничего не найдено",
		En: "Nothing found",
	},
}

var PaymentRequired = ServerError{
	Message: ResultMessage{
		Tk: "Toleg gerekli",
		Ru: "Требуется оплата",
		En: "Payment required",
	},
}

type Response struct {
	Error  error       `json:"error"`
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResultMessage struct {
	Tk string `json:"tk"`
	Ru string `json:"ru"`
	En string `json:"en"`
}

type Success struct {
	Message string `json:"message"`
}

type SuccessWithID struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type ServerError struct {
	Message ResultMessage `json:"message"`
}
