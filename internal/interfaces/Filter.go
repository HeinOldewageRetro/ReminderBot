package interfaces

type Filter interface {
	//This function must return true if the message must be sent
	Filter(Message) bool
}
