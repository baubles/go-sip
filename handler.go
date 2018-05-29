package sip

// A Handler responds to an HTTP request.
type Handler interface {
	ServeSIP(*Request) Response
}
