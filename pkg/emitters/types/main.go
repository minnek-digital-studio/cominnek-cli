package emitterTypes

type User struct {
	Name string
	Username string
}

type IRootEmitter struct {
	Ticket string
	Branch string
	Route string
	User User
}
