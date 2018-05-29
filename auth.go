package sip

// AuthManager auth manager
type AuthManager struct {
	Lookup func(realm, acc_name string) (acc *Account)

	Verify func(req *Request) bool
}
