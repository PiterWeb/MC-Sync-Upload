package src

type Users struct {
	Users []User `toml:"users"`
}

type User struct {
	Name        string
	Uuid        string `toml:omitempty`
	OfflineUuid string `toml:omitempty`
}

var users Users
