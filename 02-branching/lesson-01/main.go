package main

import "fmt"

func main() {
	score := 50
	/*
	   Вложенный код = сложно читать
	   	if score > 10 {
	   		if score > 15 {
	   			fmt.Printf("Текущий счёт: %d\nТы МЕГА-КРАСАВЧИК", score)
	   		} else {
	   			fmt.Printf("Текущий счёт: %d\nТы красавчик", score)
	   		}
	   	} else {
	   		fmt.Printf("Текущий счёт: %d\nУчись играть", score)
	   	}
	*/
	/*
		if score > 15 {
			fmt.Println("Ты МЕГА-красавчик")
		} else if score > 10 {
			fmt.Println("Ты красавчик")
		} else {
			fmt.Println("учись играть")
		}
	*/
	if score == 15 {
		fmt.Println("Дюжина!")
	} else if score == 21 {
		fmt.Println("Очко!")
	} else if score == 50 {
		fmt.Println("Полтинник!")
	} else {
		fmt.Println("Ты не попал ни в одну категорию")
	}
}
