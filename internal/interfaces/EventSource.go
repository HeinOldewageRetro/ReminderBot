package interfaces

type EventSource interface {
	//When the event source generates a message the function passed as
	//a parameter here is called with the `Message`
	Handler(func(Message))
}
