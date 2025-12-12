package main

import (
	"encoding/xml"
	"fmt"
)

type Engine struct {
	XMLName xml.Name `xml:"engine"`
	Model   string   `xml:"model,attr"`
	Power   int      `xml:"power"`
}

type Car struct {
	XMLName xml.Name `xml:"car"`
	Brand   string   `xml:"brand"`
	Model   string   `xml:"model"`
	Engine  Engine   `xml:"engine"`
}

func main() {
	car := Car{
		Brand: "Mitsubishi",
		Model: "Lancer",
		Engine: Engine{
			Model: "4B11",
			Power: 150,
		},
	}

	xmlData, err := xml.MarshalIndent(car, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	fmt.Println("Сериализованный XML:")
	fmt.Println(string(xmlData))

	var car2 Car
	err = xml.Unmarshal(xmlData, &car2)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	fmt.Printf("\nДесериализованный объект: %+v\n", car2)
}
