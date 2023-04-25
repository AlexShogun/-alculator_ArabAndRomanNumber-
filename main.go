package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Operation, OperandOne, OperandTwo := InputUser()
	// Указатели для возвращаемых значений функции InputUser
	//var Operat *string = &Operation
	//var OperOne *string = &OperandOne
	//var OperTwo *string = &OperandTwo
	for {
		Input := InputUser()
		Err := func(str []string) {
			if err := CheckErr(str); err != "" {

				fmt.Println(err)
				os.Exit(0)
			}

		}
		Err(Input)

	}

}

//Принимает ввод пользователя и возвращает строку операции, первой и второй цифры

func InputUser() (str []string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	str = strings.Split(text, " ")
	return str

}
func CheckErr(str []string) (err string) {
	var inpUser *[]string = &str      //Указатель на слайс ввода пользователя
	var lenInUser int = len(*inpUser) //Количесво элементов в слайсе
	var lenOperation int              //Количество операций в примере
	var listOperation []string = []string{"+", "-", "*", "/"}
	var system int = CheckSystem(*inpUser) //Система счисления
	var systems *int = &system             //Указатель на систему счисления
	var reply int = Сalculit(*inpUser, *systems)

	for _, i := range *inpUser {
		for _, l := range listOperation {
			if l == i {
				lenOperation += 1
			}
		}
	} // Подсчет количества операций в вводе пользователя

	if lenInUser <= 2 {
		return "Ошибка, так как строка не является математической операцией."

	} else if lenOperation > 1 {
		return "Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	} else if *systems == 3 {
		return "Ошибка, так как используются одновременно разные системы счисления."
	} else if *systems == 0 {
		return "Ошибка, так как строка не является математической операцией."
	} else if CheckNumber(*inpUser, *systems) == false {
		if *systems == 1 {
			return "Ошибка, так как формат ввода цифр не удовлетворяет заданию - от 1 до 10"
		} else if *systems == 2 {
			return "Ошибка, так как формат ввода цифр не удовлетворяет заданию - от I до X"
		}
	} else if reply <= 0 && *systems == 2 {
		return "Ошибка, так как в римской системе нет отрицательных чисел."
	}

	if *systems == 2 {
		fmt.Println(DecodArabTheRoman(reply))
	} else if *systems == 1 {
		fmt.Println(reply)
	}

	return err
}
func CheckSystem(str []string) (s int) {
	arabOne := ArabSystem(str[0])
	arabTwo := ArabSystem(str[2])
	romanOne := RomanSystem(str[0])
	romanTwo := RomanSystem(str[2])

	if arabOne != 0 && arabTwo != 0 {
		return 1 // Арабская
	} else if romanOne != 0 && romanTwo != 0 {
		return 2 //Римская
	} else if (arabOne != 0 && romanTwo != 0) || (romanOne != 0 && arabTwo != 0) {
		return 3 // Смешанная система (ошибка)
	}

	return s
}
func ArabSystem(a string) int {
	one, err := strconv.Atoi(a)
	if err != nil {

	}
	return one

}
func RomanSystem(s string) (result int) {

	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result = 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result

}
func CheckNumber(str []string, s int) (t bool) {
	var a int
	var b int
	// Арабская система
	if s == 1 {
		a = ArabSystem(str[0])
		b = ArabSystem(str[2])
	} else if s == 2 {
		a = RomanSystem(str[0])
		b = RomanSystem(str[2])
	}
	if (1 <= a && a <= 10) && (1 <= b && b <= 10) {
		return true
	}
	return t
}
func Сalculit(str []string, s int) (reply int) {
	system := func(str []string, s int) (a, b int) {

		// Арабская система
		if s == 1 {
			a = ArabSystem(str[0])
			b = ArabSystem(str[2])
		} else if s == 2 { //Римская система
			a = RomanSystem(str[0])
			b = RomanSystem(str[2])
		}
		return a, b
	}
	a, b := system(str, s)
	var Operation *string = &str[1]
	switch *Operation {
	case "+":
		reply = a + b
	case "-":
		reply = a - b
	case "*":
		reply = a * b
	case "/":
		reply = a / b
	}
	return reply
}
func DecodArabTheRoman(meaning int) (reply string) {
	var romanNumbers map[int]string = map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	var keys []int = []int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var numbers []int

	for meaning != 0 {
		for _, key := range keys {
			if key == meaning {
				meaning -= key
				numbers = append(numbers, key)
			} else if key < meaning {
				meaning -= key
				numbers = append(numbers, key)
				break
			}
		}
	}

	var roman string

	for _, keys := range numbers {
		for clue, value := range romanNumbers {
			if clue != keys {
				continue
			} else if clue == keys {
				roman += value
			}
		}
	}

	reply += roman
	return reply
}
