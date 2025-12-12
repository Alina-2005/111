package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Engine struct {
	Model string `json:"model"`
	Power int    `json:"power_hp"`
}

type Car struct {
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Engine Engine `json:"engine"`
}

func main() {
	car := Car{
		Brand: "Hyundai",
		Model: "Solaris",
		Engine: Engine{
			Model: "Gamma",
			Power: 123,
		},
	}

	// Сериализация (marshal)
	jsonData, err := json.MarshalIndent(car, "", "  ")
	if err != nil {
		log.Fatal("Ошибка сериализации:", err)
	}
	fmt.Println("Сериализованный JSON:")
	fmt.Println(string(jsonData))

	// Десериализация (unmarshal)
	var car2 Car
	err = json.Unmarshal(jsonData, &car2)
	if err != nil {
		log.Fatal("Ошибка десериализации:", err)
	}
	fmt.Printf("\nДесериализованный объект: %+v\n", car2)
}
