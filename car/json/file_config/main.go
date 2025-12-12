package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type EngineConfig struct {
	Type   string `json:"type"`
	Volume string `json:"volume,omitempty"`
}

type Car struct {
	Model      string       `json:"model"`
	Year       int          `json:"year"`
	Color      string       `json:"color"`
	Engine     EngineConfig `json:"engine"`
	HasSunroof bool         `json:"has_sunroof"`
}

func main() {
	config := Car{
		Model: "Renault Logan",
		Year:  2019,
		Color: "White",
		Engine: EngineConfig{
			Type:   "K7M",
			Volume: "1.6L",
		},
		HasSunroof: false,
	}

	// Запись в файл
	filePath := "car.json"
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Не удалось создать файл:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err = encoder.Encode(config); err != nil {
		log.Fatal("Ошибка записи в файл:", err)
	}

	fmt.Printf("Конфигурация автомобиля сохранена в %s\n", filePath)

	// Чтение из файла
	file2, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Не удалось открыть файл:", err)
	}
	defer file2.Close()

	var loaded Car
	if err = json.NewDecoder(file2).Decode(&loaded); err != nil {
		log.Fatal("Ошибка чтения из файла:", err)
	}

	fmt.Printf("Загружено из файла: %+v\n", loaded)
}
