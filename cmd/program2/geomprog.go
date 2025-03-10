package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func sliceGeneration() []int {
	length := 5 + rand.Intn(5)
	var valueSlice []int
	k := 2 + rand.Intn(5)
	multiplicator := k
	for i := 0; i < length; i++ {
		valueSlice = append(valueSlice, k)
		k *= multiplicator
	}

	return valueSlice
}

func taskGeneration(name string) bool {
	taskSlice := sliceGeneration()

	k := rand.Intn(len(taskSlice))

	var output string

	for i, v := range taskSlice {
		if i == k {
			output += ".. "
			continue
		}

		output += strconv.Itoa(v) + " "
	}

	var answer int
	fmt.Printf("Question: %s\nYour answer: ", output)
	_, err := fmt.Scan(&answer)

	if err != nil {
		log.Fatal(err)
	}

	if answer == taskSlice[k] {
		fmt.Printf("Correct!\n"+"Congratulations, %s!\n", name)
		return true
	} else {
		fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'\n", answer, taskSlice[k])
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
		if !taskGeneration(name) {
			break
		}
	}

}
