package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

// Цвета для терминала (ANSI-escape коды)
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

// GameResult описывает структуру записи для results.json
type GameResult struct {
	Date     string `json:"date"`
	Outcome  string `json:"outcome"`
	Attempts int    `json:"attempts"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		maxNum, tryCount, err := chooseDifficulty(reader)
		if err != nil {
			fmt.Println("\nВвод завершён. До встречи! 👋")
			return
		}

		winNum := genWinNum(maxNum)
		var history []int
		isWin := false

		fmt.Println("\n🦆💨 Игра началась! ✨")
		fmt.Printf("🎰 Я загадал число от 1 до %d! У тебя %s%d попыток%s чтобы угадать его 🎲\n\n", maxNum, colorYellow, tryCount, colorReset)

		for tries := 1; tries <= tryCount; tries++ {
			fmt.Printf("💫 Итак, %sпопытка номер %d%s (осталось: %s%d%s), твоё число? 💫\n",
				colorYellow, tries, colorReset,
				colorYellow, tryCount-tries+1, colorReset)

			userNum, err := userEnter(reader)
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\nВвод завершён. До встречи! 👋")
					return
				}

				fmt.Println("Нужно ввести целое число 🤓")
				tries-- //откат попытки
				continue
			}

			history = append(history, userNum) // записываем в слайс введённое число

			if checkGuess(userNum, winNum) {
				isWin = true
				saveResult("Победа", tries)
				break
			}

			giveHint(userNum, winNum)
			fmt.Printf("Твои попытки: %s%v%s\n\n", colorYellow, history, colorReset)
		}

		if !isWin {
			fmt.Printf("\n%s🦆💨 Попытки кончились, ты не угадал(( число было: %d%s\n", colorRed, winNum, colorReset)
			saveResult("Поражение", tryCount)
		}

		if !askPlayAgain(reader) {
			fmt.Println("Спасибо за игру! До встречи 👋")
			break
		}
		fmt.Println()
	}
}

// chooseDifficulty отвечает за выбор уровня сложности (Фича 4)
func chooseDifficulty(reader *bufio.Reader) (int, int, error) {
	for {
		fmt.Println("Выбери уровень сложности:")
		fmt.Println("1 — Easy (числа 1–50, 15 попыток)")
		fmt.Println("2 — Medium (числа 1–100, 10 попыток)")
		fmt.Println("3 — Hard (числа 1–200, 5 попыток)")
		fmt.Print("Твой выбор (1/2/3): ")

		choice, err := userEnter(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return 0, 0, err
			}

			fmt.Println("Ошибка! Введи цифру 1, 2 или 3.\n")
			continue
		}

		switch choice {
		case 1:
			return 50, 15, nil
		case 2:
			return 100, 10, nil
		case 3:
			return 200, 5, nil
		default:
			fmt.Println("Нет такого варианта. Введи 1, 2 или 3.\n")
		}
	}
}

// genWinNum генерирует случайное число от 1 до maxNum
func genWinNum(maxNum int) int {
	return rand.IntN(maxNum) + 1
}

// userEnter считывает и проверяет ввод пользователя
func userEnter(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

// checkGuess сравнивает число игрока с загаданным
func checkGuess(userNum, winNum int) bool {
	if userNum == winNum {
		fmt.Printf("%s💃🥁 Ого, ты угадал! 🎈👑%s\n", colorGreen, colorReset)
		return true
	} else if userNum < winNum {
		fmt.Println("Ноу, моё число больше 👆")
	} else {
		fmt.Println("Ноу, моё число меньше 👇")
	}
	return false
}

// giveHint выводит подсказку по расстоянию до загаданного числа
func giveHint(userNum, winNum int) {
	diff := winNum - userNum
	if diff < 0 {
		diff = -diff
	}

	if diff <= 5 {
		fmt.Println(" 🌋 Горячо! 🌡🔥")
	} else if diff <= 15 {
		fmt.Println(" 🌄 Тепло! 🌤🐝")
	} else {
		fmt.Println(" 🗻 Холодно! ⛄❄")
	}
}

// askPlayAgain спрашивает, хочет ли игрок сыграть ещё раз (Фича 5)
func askPlayAgain(reader *bufio.Reader) bool {
	for {
		fmt.Print("\nХочешь сыграть ещё раз? (д/н или y/n): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return false
		}
		answer := strings.ToLower(strings.TrimSpace(input))

		if answer == "д" || answer == "да" || answer == "y" || answer == "yes" {
			return true
		}
		if answer == "н" || answer == "нет" || answer == "n" || answer == "no" {
			return false
		}
		fmt.Println("Пожалуйста, введи 'д' (да) или 'н' (нет).")
	}
}

// saveResult сохраняет результат игры в файл results.json (Фича 6)
func saveResult(outcome string, attempts int) {
	const filename = "results.json"
	var results []GameResult

	// Читаем существующий файл, если он есть
	data, err := os.ReadFile(filename)
	if err == nil && len(data) > 0 {
		_ = json.Unmarshal(data, &results)
	}

	newRecord := GameResult{
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Outcome:  outcome,
		Attempts: attempts,
	}
	results = append(results, newRecord)

	updatedData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("Не удалось подготовить данные для сохранения:", err)
		return
	}

	if err := os.WriteFile(filename, updatedData, 0644); err != nil {
		fmt.Println("Не удалось сохранить результат в файл:", err)
	}
}
