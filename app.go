package main

import "fmt"

func main() {
	result := pension("mae", 78)
fmt.Println(result)
}

func pension(gender string, age int) string {
	if gender == "male" && age >= 63 {
		return "Пора на пенсию"
	} else if gender == "female" && age >= 61 {
		return "Пора на пенсию"
	} else if gender != "male" && gender != "female" {
		return "Ты чо пишешь, чмо"
	}
	return "Работай, сынок"
}
