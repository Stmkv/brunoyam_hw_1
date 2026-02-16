package main

import (
	"fmt"
	"homework/internal/input"
	"unicode"
)

func main() {
	fakeDb := []string{
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
			names := searchName(&fakeDb, userMenuInput)
			fmt.Println(names)
		}
		if userMenuItem == "2" {
			fmt.Println("Введите Новое имя ")
			userMenuInput := input.ReadLine()
			addPerson(&fakeDb, userMenuInput)
			fmt.Println(fakeDb)

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

func searchName(fDb *[]string, firstLetterName string) []string {
	userFirstLetterNameRune := []rune(firstLetterName)[0]
	correctFirstLetterUser := unicode.ToLower(userFirstLetterNameRune)

	var namesFound = make([]string, 0)
	for _, name := range *fDb {
		nameRune := unicode.ToLower([]rune(name)[0])
		if nameRune == correctFirstLetterUser {
			namesFound = append(namesFound, name)
		}
	}
	return namesFound
}

func addPerson(fDb *[]string, userName string) {
	*fDb = append(*fDb, userName)
}
