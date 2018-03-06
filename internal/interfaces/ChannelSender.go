package interfaces

type ChannelSender interface {
	Send(Message) error
}
