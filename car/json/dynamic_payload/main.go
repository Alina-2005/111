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
	// Динамический JSON (неизвестная структура)
	dynamicJSON := `{
		"type": "vehicle_info",
		"payload": {
			"brand": "Kia",
			"model": "Rio",
			"engine": {
				"model": "Kappa",
				"power_hp": 100
			}
		}
	}`

	type Event struct {
		Type    string          `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}

	var event Event
	if err := json.Unmarshal([]byte(dynamicJSON), &event); err != nil {
		log.Fatal("Ошибка парсинга в Event:", err)
	}

	fmt.Printf("Тип сообщения: %s\n", event.Type)

	if event.Type == "vehicle_info" {
		var car Car
		if err := json.Unmarshal(event.Payload, &car); err != nil {
			log.Fatal("Ошибка парсинга Payload в Car:", err)
		}
		fmt.Printf("Распакованные данные автомобиля: %+v\n", car)
		fmt.Printf("Двигатель: %s, Мощность: %d л.с.\n", car.Engine.Model, car.Engine.Power)
	}
}
