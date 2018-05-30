package sip

type Cred struct {
	Username string
	Realm    string // Realm. Use "*" to make a credential that can be used to authenticate against any challenges.
	Scheme   string // Scheme (e.g. "digest").
	DataType int    //Type of data (0 for plaintext passwd).
	Data     string // The data, which can be a plaintext password or a hashed digest.
}
