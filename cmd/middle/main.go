package main

import (
	"fmt"
	"homework/internal/input"
	"unicode"
)

func main() {
	fakeDB := []string{
		"Анна",
		"Борис",
		"Валентин",
		"Степан",
		"Михаил",
		"Алена",
		"Полина",
	}
	for {
		printMainMenu()
		userMenuItem := input.ReadLine()
		if userMenuItem == "1" {
			fmt.Println("Введите первую букву имени ")
			userMenuInput := input.ReadLine()
			names := searchName(&fakeDB, userMenuInput)
			fmt.Println(names)
		}
		if userMenuItem == "2" {
			fmt.Println("Введите Новое имя ")
			userMenuInput := input.ReadLine()
			addPerson(&fakeDB, userMenuInput)
			fmt.Println(fakeDB)
		}
		if userMenuItem == "3" {
			break
		}
	}
}

func printMainMenu() {
	fmt.Println("Выберите пункт, чтобы начать. ")
	fmt.Println(`1. Найти все имена начинающиеся с буквы
2. Добавить новое имя 
3. Завершить работу `)
}

func searchName(fDB *[]string, firstLetterName string) []string {
	userFirstLetterNameRune := []rune(firstLetterName)[0]
	correctFirstLetterUser := unicode.ToLower(userFirstLetterNameRune)

	var namesFound = make([]string, 0)
	for _, name := range *fDB {
		nameRune := unicode.ToLower([]rune(name)[0])
		if nameRune == correctFirstLetterUser {
			namesFound = append(namesFound, name)
		}
	}
	return namesFound
}

func addPerson(fDB *[]string, userName string) {
	*fDB = append(*fDB, userName)
}
