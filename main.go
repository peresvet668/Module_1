package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {

	tryCount := 10 //количество попыток
	maxNum := 100  //максимальное число

	winNum := genWinNum(maxNum)

	reader := bufio.NewReader(os.Stdin)

	var history []int

	fmt.Println("🦆💨 Игра началась! ✨")
	fmt.Printf("🎰 Я загадал число от 1 до %d! У тебя %d попыток чтобы угадать его 🎲\n", maxNum, tryCount)

	for tries := 1; tries <= tryCount; tries++ {
		fmt.Printf("💫 Итак, попытка номер %d, твоё число? 💫\n", tries)

		userNum, err := userEnter(reader)
		if err != nil {
			fmt.Println("Нужно ввести целое число 🤓")
			tries-- // откат попытки
			continue
		}

		history = append(history, userNum) //записываем в слайс введённое число

		if userNum == winNum { // Сравниваем введённое с загаданным
			fmt.Println("💃🥁 Ого, ты угадал! 🎈👑")
			return
		} else if userNum < winNum {
			fmt.Println("Ноу, моё число больше 👆")
		} else {
			fmt.Println("Ноу, моё число меньше 👇")
		}

		diff := winNum - userNum //вычисляем насколько далеко значение

		if diff < 0 {
			diff = -diff //если число больше, то нам надо избавиться от минуса
		}

		if diff <= 5 {
			fmt.Println(" 🌋 Горячо! 🌡🔥")
		} else if diff <= 15 {
			fmt.Println(" 🌄 Тепло! 🌤🐝")
		} else {
			fmt.Println(" 🗻 Холодно! ⛄❄")
		}

		fmt.Printf("Твои попытки: %v\n\n", history)
	}

	fmt.Printf("\n 🦆💨 Попытки кончились, ты не угадал(( число было: %d\n", winNum)

}

func genWinNum(maxNum int) int {
	return rand.IntN(maxNum) + 1
}

func userEnter(reader *bufio.Reader) (int, error) { //проверка ввода
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}
