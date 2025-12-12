package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Engine struct {
	Model string
	Power int
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine
}

func main() {
	car := Car{
		Brand: "Ford",
		Model: "Focus",
		Engine: Engine{
			Model: "Duratec",
			Power: 125,
		},
	}

	// Буфер для хранения данных
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(car)
	if err != nil {
		log.Fatal("Ошибка gob.Encode:", err)
	}

	// Десериализация
	var car2 Car
	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&car2)
	if err != nil {
		log.Fatal("Ошибка gob.Decode:", err)
	}

	fmt.Printf("Исходный автомобиль: %+v\n", car)
	fmt.Printf("Восстановленный автомобиль: %+v\n", car2)

	// Проверим, что данные идентичны
	if car.Brand == car2.Brand && car.Model == car2.Model && car.Engine.Model == car2.Engine.Model {
		fmt.Println("Gob сериализация/десериализация прошла успешно!")
	}
}
