package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*for i := 1; i <= 3; i++ {
		fmt.Println("Hello")
	}
	*/
	//	fmt.Println("Случайное число от 0 до 9:", rand.Intn(10)) //генерация случайных чисел в заданном диапазоне rand.Intn(10) от 0 до 9

	fmt.Println("Бросок монетки:")
	fmt.Println("")

	if rand.Intn(2) == 1 { //при таком примитивном применении == 1 в rand.Intn(x), увеличивая число x мы уменьшаем вероятность выполнения
		fmt.Println("Выпал ОРЁЛ")
	} else {
		fmt.Println("Выпала РЕЖКА")
	}
}
