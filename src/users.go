package src

type Users struct {
	Users []User `toml:"users"`
}

type User struct {
	Name        string
	Uuid        string
	OfflineUuid string
}

var users Users