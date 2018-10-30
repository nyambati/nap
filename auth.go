package nap

// AuthToken struc
type AuthToken struct {
	Token string
}

// AuthBasic struc
type AuthBasic struct {
	Username string
	Password string
}

// Authentication struct
type Authentication interface {
	AuthorizationHeader() string
}
