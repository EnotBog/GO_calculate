package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkNumber(buf string) bool {
	for _, v := range buf {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

func resultSymbol2(buf string) int {
	var res int
	buf_up := 0
	var buf_res int
	var pair_rim = map[rune]int{'X': 10, 'V': 5, 'I': 1}
	//если следующее число больше, то минусуем из большего если равно добавляем в буфер, если меньше прибавляем в res
	//
	for k := 0; k < len(buf); k++ {
		if pair_rim[rune(buf[k])] != 0 {
			// свод правил старшая цифра вычитаем меньшую
			for {
				if buf_up <= pair_rim[rune(buf[k])] {
					buf_up = pair_rim[rune(buf[k])]
					if k != 0 && pair_rim[rune(buf[k-1])] < pair_rim[rune(buf[k])] {
						buf_res = pair_rim[rune(buf[k])] - buf_res
						res = res + buf_res
						buf_res = 0
					} else if k == 0 && len(buf) == 1 {
						buf_res = buf_res + pair_rim[rune(buf[k])]
						res = res + buf_res
						buf_res = 0
						break
					} else if buf_up < pair_rim[rune(buf[k])] || k == 0 && (pair_rim[rune(buf[k])] < pair_rim[rune(buf[k+1])]) {
						buf_res = pair_rim[rune(buf[k])]
					} else {
						buf_res = buf_res + pair_rim[rune(buf[k])]
						res = res + buf_res
						buf_res = 0
					}
					break
				} else if buf_up >= pair_rim[rune(buf[k])] {
					if k == len(buf)-1 {
						res = res + pair_rim[rune(buf[k])] + buf_res
					} else {
						buf_res = pair_rim[rune(buf[k])]
						buf_up = pair_rim[rune(buf[k])]
					}
					break
				}
			}
		}
	}
	if res == 0 {
		return -1
	}
	return res
}

func FindOper(buf string) (bool, int) {
	res_bool := false
	res_index := -1
	symbol_oper := []rune{'+', '-', '*', '/'}
	count := 0
	for i := range symbol_oper {
		count += strings.Count(buf, string(symbol_oper[i]))
		if count > 1 {
			panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		}
	}

	for i := range symbol_oper {
		if strings.IndexRune(buf, symbol_oper[i]) != -1 {
			res_index = strings.IndexRune(buf, symbol_oper[i])
			res_bool = true
			count++
		}
	}

	if res_index == -1 {
		res_bool = false
	}
	return res_bool, res_index
}

func ReadString(buf string) (string, string, string) {
	if len(buf) == 0 {
		panic("Отсутствуют данные")
	}
	var a, b, c string

	if bool_check, index := FindOper(buf); bool_check { // берем индекс пердолим его как разделитель
		b = string(buf[index])
		a = buf[0:index]
		c = buf[index+1:]
	} else {
		panic("Нет знака действия")
	}
	return a, b, c
}

func FuncOperation(buf string) (int, bool) {
	var a_buf, c_buf int
	string_forma := strings.Trim(buf, " ")
	a, b, c := ReadString(string_forma)
	var arab_number bool = false
	if checkNumber(a) && checkNumber(c) {
		a_buf, _ = strconv.Atoi(a)
		c_buf, _ = strconv.Atoi(c)
		arab_number = true
	} else if !checkNumber(a) && !checkNumber(c) {
		a_buf = resultSymbol2(a)
		c_buf = resultSymbol2(c)
		if a_buf == -1 {
			log.Fatal("Число ", a, " неправильное")
		}
		if c_buf == -1 {
			log.Fatal("Число ", c, " неправильное")
		}
		if c_buf > 10 || a_buf > 10 {
			panic("Одно из введенных чисел число больше 10")
		}
	} else {
		panic("Действие с арабским и римским числом невозможно  ")
	}
	switch b {
	case "+":
		return a_buf + c_buf, arab_number
	case "-":
		return a_buf - c_buf, arab_number
	case "*":
		return a_buf * c_buf, arab_number
	case "/":
		if c_buf == 0 {
			panic("Делитель 0")
		}
		return a_buf / c_buf, arab_number
	}
	return -128, arab_number
}
func Result(buf string) string {
	if res, arab_number := FuncOperation(buf); arab_number {
		return strconv.Itoa(res)
	} else {
		if res <= 0 {
			panic("Результат меньше 0")
		}
		return ConvertR(res)
	}
}

func ConvertR(res int) string {
	var string_res string
	var pair_rim = map[int]string{10: "X", 9: "IX", 8: "VIII", 7: "VII", 6: "VI", 5: "V", 4: "IV", 3: "III", 2: "II", 1: "I"}
	for res != 0 {
		if res >= 10 {
			string_res += "X"
			res -= 10
		} else {
			switch res {
			case 9:
				string_res += pair_rim[9]
				res -= 9
			case 8:
				string_res += pair_rim[8]
				res -= 8
			case 7:
				string_res += pair_rim[7]
				res -= 7
			case 6:
				string_res += pair_rim[6]
				res -= 6
			case 5:
				string_res += pair_rim[5]
				res -= 5
			case 4:
				string_res += pair_rim[4]
				res -= 4
			case 3:
				string_res += pair_rim[3]
				res -= 3
			case 2:
				res -= 2
				string_res += pair_rim[2]
			case 1:
				res -= 1
				string_res += pair_rim[1]
			}
		}
	}
	return string_res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример")
	in_string, _ := reader.ReadString('\n')

	fmt.Println(Result(in_string))

}
