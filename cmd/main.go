package main

import "student-crud/internal/app"

func main() {
	// запуск приложения
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
