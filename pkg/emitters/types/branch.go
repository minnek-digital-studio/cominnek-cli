package emitterTypes

type IBranchEventData struct {
	Branch string
	Ticket string
	Type string
}

type IBranchFailedData struct {
	Error string
	Data IBranchEventData
}