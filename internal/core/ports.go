package core

type FormLogic interface {
	SetFields(message string)
	Log()
	Submit() error
}

type Logic interface {
	FormLogic
}

type AppState struct {
	message string
}
