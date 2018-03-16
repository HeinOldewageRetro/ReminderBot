package interfaces

type Message interface {
	//So we need some way to get the message
	//And the things to filter on (Target audience,Country)
	String() string
	Target() []string
	Country() string
}
