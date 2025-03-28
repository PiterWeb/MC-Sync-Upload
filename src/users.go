package src

type Users struct {
	Users []User `toml:"users"`
}

type User struct {
	Name        string
	Uuid        string `toml:"uuid,omitempty"`
	OfflineUuid string `toml:"offlineUuid,omitempty"`
}

var users Users
