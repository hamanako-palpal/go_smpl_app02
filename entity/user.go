package entity

type User struct {
	Name string
	Pass string
}

type AdminUser struct {
	User
	AddUser func(nm string, ps string)
}
