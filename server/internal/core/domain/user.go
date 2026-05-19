package domain

type User struct {
	ID      string
	Version int

	FullName    string
	PhoneNumber *string
}
