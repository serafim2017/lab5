package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем срез с разными приветствиями
	greetings := []string{
		"Chao!",
		"Hola!",
		"Привет!",
		"Hello",
	}

	// Выбираем случайное приветствие из среза
	randomIndex := rand.Intn(len(greetings))
	randomGreeting := greetings[randomIndex]

	// Выводим случайное приветствие
	fmt.Println(randomGreeting)
}
