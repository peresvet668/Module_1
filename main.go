package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var tryCount int //количество попыток
var maxNum int   //максимальное число

func main() {

	tryCount = 10
	maxNum = 100

	winNum := genWinNum(maxNum)

	reader := bufio.NewReader(os.Stdin)

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

		if userNum == winNum { // Сравниваем введённое с загаданным
			fmt.Println("💃🥁 Ого, ты угадал! 🎈👑")
			return
		} else if userNum < winNum {
			fmt.Println("Ноу, моё число больше 👆")
		} else {
			fmt.Println("Ноу, моё число меньше 👇")
		}
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
