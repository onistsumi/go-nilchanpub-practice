package main

import "fmt"

func main() {
	officianWork("Андрей")
	officianWork("Вася")
}

func officianWork(name string) {
	fmt.Println("Накрываю на стол")
	fmt.Println("Привет,", name, "!")
	fmt.Println("Принимаю заказ...")
	fmt.Println("Принёс заказ!")
	fmt.Println("")
}

func hello() {
	fmt.Println("Я функция hello")
	fmt.Println("Меня вызвали через hello()")
	fmt.Println("Я завершаюсь")
	fmt.Println("")
}

func square(x int) {
	fmt.Println("Мы приняли в функцию параметр x:", x)
	fmt.Println("Квадрат х равен:", x*x)
}
