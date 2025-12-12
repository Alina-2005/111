package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Запуск всех примеров сериализации автомобилей...")

	baseDir := "car"

	exampleDirs := []string{
		"gob",
		"json/simple",
		"json/custom_date",
		"json/dynamic_payload",
		"json/file_config",
		"xml",
	}

	for _, dir := range exampleDirs {
		path := "./" + filepath.ToSlash(filepath.Join(baseDir, dir))
		fmt.Printf("\n--- Запуск примера из: %s ---\n", path)

		cmd := exec.Command("go", "run", path)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Ошибка при выполнении %s:\n", path)
			fmt.Println(string(output))
		} else {
			fmt.Println(string(output))
		}
		fmt.Printf("--- Завершено: %s ---\n", path)
	}
}
