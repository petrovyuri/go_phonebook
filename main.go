package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *[]Entry {
	result := make([]Entry, 0)
	for _, v := range data {
		if v.Name == key {
			result = append(result, v)
		}
	}
	if len(result) > 0 {
		return &result
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	data = append(data, Entry{"Yura", "Petrov", "2109416471"})
	data = append(data, Entry{"Sveta", "Volkova", "2109416871"})
	data = append(data, Entry{"Oleg", "Fin", "2109416123"})
	data = append(data, Entry{"Yura", "Ivanova", "2109416123"})

	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Println(exe, "поиск|список <arguments>")
		return
	}

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Используй: search Yura")
			return
		}

		result := search(arguments[2])

		if result == nil {
			fmt.Println("Не найдено:", arguments[2])
			return
		}

		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Команда не найдена")
	}
}
