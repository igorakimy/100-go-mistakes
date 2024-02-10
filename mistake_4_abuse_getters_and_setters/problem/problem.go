package main

// Избыточное использование геттеров и сеттеров.
// Можно обойтись без них, объявив публичными свойства структуры

type User struct {
	firstName string
	lastName  string
	age       int
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) SetAge(age int) {
	u.age = age
}
