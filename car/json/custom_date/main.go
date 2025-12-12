package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type CustomDate struct {
	time.Time
}

// CustomTime сериализуется как строка в формате YYYY-MM-DD
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Format("2006-01-02"))
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

type Car struct {
	Model          string     `json:"model"`
	ProductionDate CustomDate `json:"production_date"`
}

func main() {
	prodDate := time.Date(2021, 8, 10, 0, 0, 0, 0, time.UTC)
	car := Car{
		Model:          "Lada Vesta",
		ProductionDate: CustomDate{prodDate},
	}

	// Сериализация
	jsonData, err := json.Marshal(car)
	if err != nil {
		log.Fatal("Ошибка сериализации:", err)
	}
	fmt.Println("Сериализованный JSON:", string(jsonData))

	// Десериализация
	var car2 Car
	err = json.Unmarshal(jsonData, &car2)
	if err != nil {
		log.Fatal("Ошибка десериализации:", err)
	}
	fmt.Printf("Десериализованный объект: Model=%s, ProductionDate=%s\n",
		car2.Model,
		car2.ProductionDate.Format("2006-01-02"))

	// Проверка равенства
	if car.Model == car2.Model && car.ProductionDate.Equal(car2.ProductionDate.Time) {
		fmt.Println("Кастомная сериализация для автомобиля работает!")
	}
}
