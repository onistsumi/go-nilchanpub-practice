package main

import "fmt"

type User struct { // это описание структуры User, банально то, что будет входить в структуру User
	name           string
	age            int
	phoneNumber    string
	isCloseAccaunt bool
	rating         float64
}

func main() {
	user1 := User{ // а вот это уже СОЗДАНИЕ нового пользователя по шаблону User
		name:           "Джон",
		age:            37,
		phoneNumber:    "+0 8392-0234",
		isCloseAccaunt: true,
		rating:         7.2,
	}
	/*
		fmt.Println("Пользователь №1:", user1)
		fmt.Println("Изначальное имя пользователя №1:", user1.name)
		user1.name = "Васёк"
		fmt.Println("Изменённое имя пользователя №1:", user1.name)
	*/
	Greeting(user1)
}

func Greeting(u User) {
	fmt.Println("Всем привет :)")
	fmt.Println("Меня зовут:", u.name)
	fmt.Println("Мой рейтинг:", u.rating)
}
