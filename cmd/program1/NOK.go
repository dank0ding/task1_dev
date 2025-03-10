package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

// Запись всех возможных вариантов кратного для каждого элемента
func checkMultiplicity(number int) map[int]int {
	checkMap := make(map[int]int)
	addVar := number
	checkMap[number] = 1

	for i := 0; i < 100; i++ {
		number += addVar
		checkMap[number] = 1
	}

	return checkMap
}

// Поиск наименьшего общего кратного
func findNumber(a, b, c map[int]int) int {

	var keySlice []int
	for k := range a {
		if a[k] == 1 && b[k] == 1 && c[k] == 1 {
			keySlice = append(keySlice, k)
		}
	}

	sort.Ints(keySlice)

	return keySlice[0]
}


// Проверка ответа пользователя
func checkAnswer(answer, checkAnswer int) bool {
	if answer == checkAnswer {
		return true
	} else {
		return false
	}

}

// Основной алгоритм работы
func gameMain(name string) bool {

	value1 := 1 + rand.Intn(10)
	value2 := 1 + rand.Intn(10)
	value3 := 1 + rand.Intn(10)

	fmt.Printf("Question: %d %d %d", value1, value2, value3)
	var answer int
	fmt.Print("\nYour answer: ")
	_, err := fmt.Scan(&answer)

	if err != nil {
		log.Fatal(err)
	}

	mapOne := checkMultiplicity(value1)
	mapTwo := checkMultiplicity(value2)
	mapThree := checkMultiplicity(value3)

	correctAns := findNumber(mapOne, mapTwo, mapThree)

	if checkAnswer(answer, correctAns) {
		fmt.Printf("Correct!\n"+"Congratulations, %s!\n", name)
		return true
	} else {
		fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'\n", answer, correctAns)
		fmt.Printf("Let's try again, %s!", name)
		return false
	}
}

func main() {
	var name string
	fmt.Println("Welcome to the Brain Games!\nMay I have your name?")
	_, err := fmt.Scan(&name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello, %s!\nFind the smallest compilation of given numbers.\n", name)

	for {
		if !gameMain(name) {
			break
		}
	}
}
