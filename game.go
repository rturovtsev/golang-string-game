package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	time2 "time"
)

func main() {
	result := false

	fmt.Println("Игра Угадай число. Есть 5 попыток.")

	time := time2.Now().Unix()
	rand.Seed(time)
	number := rand.Intn(100) + 1
	fmt.Println("загадано число от 1 до 100, угадывай!")
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 8; i++ {
		fmt.Print("Ваш вариант: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		userNumber, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if userNumber == number {
			result = true
			fmt.Println("Победа")
			break
		} else if userNumber < number {
			fmt.Println("Вы число меньше загаданного. Осталось", 6-i, "попыток")
		} else {
			fmt.Println("Ваши число больше загаданного. Осталось", 6-i, "попыток")
		}
	}

	if !result {
		fmt.Println("Ваше проиграли, попытки кончились")
	}
}
