package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func zadacha1_1() {
	var (
		value                string
		base, base2, decimal int
		decimalToBase2       []string
	)
	charToValue := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
		'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
		'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35,
	}

	decimalToBase36 := map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
		10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F", 16: "G", 17: "H", 18: "I", 19: "J",
		20: "K", 21: "L", 22: "M", 23: "N", 24: "O", 25: "P", 26: "Q", 27: "R", 28: "S", 29: "T",
		30: "U", 31: "V", 32: "W", 33: "X", 34: "Y", 35: "Z",
	}

	fmt.Print("Введите число, систему счисления и ситему счисления в которую желаете его перевезти:")
	fmt.Scan(&value, &base, &base2)
	for i, char := range value {
		exponent := len(value) - 1 - i
		decimal += charToValue[char] * int(math.Pow(float64(base), float64(exponent)))
	}
	decimal2 := decimal
	for decimal2 > 0 {
		decimalToBase2 = append(decimalToBase2, decimalToBase36[decimal2%base2])
		decimal2 = decimal2 / base2
	}
	decimalToBase2 = reverseArray(decimalToBase2)
	rezult := strings.Join(decimalToBase2, "")
	fmt.Print(rezult, "\n")

}

func zadacha1_2() {
	var (
		a, b, c, D, x1, x2 float64
		x3, x4             complex64
	)

	fmt.Print("Введите коэффициенты через пробелы: ")
	fmt.Scan(&a, &b, &c)

	D = float64(b*b - 4*a*c)

	if D > 0 {
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Printf("x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if D == 0 {
		x1 = -b / (2 * a)
		fmt.Printf("x = %.2f\n", x1)
	} else {
		imaginaryPart := math.Sqrt(-D) / (2 * float64(a))
		realPart := -b / (2 * a)
		x3 = complex(float32(realPart), float32(imaginaryPart))
		x4 = complex(float32(realPart), -float32(imaginaryPart))
		fmt.Printf("Комплексные корни: x1 = %.2f + %.2fi, x2 = %.2f %.2fi\n", real(x3), imag(x3), real(x4), imag(x4))
	}
}

func zadacha1_3() {
	var (
		text string
		num  []int
	)
	fmt.Print("Введите значения массива через запятую:")
	fmt.Scan(&text)
	numtxt := strings.Split(text, ",")
	for _, r := range numtxt {
		digit, _ := strconv.Atoi(r)
		num = append(num, digit)
	}
	sort.Slice(num, func(i, j int) bool {
		return math.Abs(float64(num[i])) < math.Abs(float64(num[j]))
	})
	fmt.Print(num, "\n")
}

func zadacha1_4() {
	var (
		text       string
		arr1, arr2 []int
	)
	fmt.Print("Введите значения первого массива через запятую:")
	fmt.Scan(&text)
	numtxt := strings.Split(text, ",")
	for _, r := range numtxt {
		digit, _ := strconv.Atoi(r)
		arr1 = append(arr1, digit)
	}
	fmt.Print("Введите значения второго массива через запятую:")
	fmt.Scan(&text)
	numtxt = strings.Split(text, ",")
	for _, r := range numtxt {
		digit, _ := strconv.Atoi(r)
		arr2 = append(arr2, digit)
	}
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			k++
			i++
		} else {
			merged[k] = arr2[j]
			k++
			j++
		}
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		k++
		j++
	}
	fmt.Print(merged, "\n")
}

func zadacha1_5() {
	var (
		text1, text2 string
	)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите первую строку: ")
	text1, _ = reader.ReadString('\n')
	text1 = strings.TrimSpace(text1)

	fmt.Print("Введите вторую строку: ")
	text2, _ = reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	n, m := len(text1), len(text2)
	if m > n {
		fmt.Println("-1")
		return
	}
	for i, j := 0, 0; i <= n-m; i++ {
		for j < m && text1[i+j] == text2[j] {
			j++
		}
		if j == m {
			fmt.Printf("Индес вхождения: %d\n", i)
			return
		}
	}
	fmt.Println("-1")
}

func zadacha2_1() {
	fmt.Print("Введите выражение через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) != 3 {
		fmt.Println("Ошибка: ожидается два числа и оператор.")
		return
	}

	a, err1 := strconv.ParseFloat(parts[0], 64)
	b, err2 := strconv.ParseFloat(parts[2], 64)
	oper := parts[1]

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка: неверный ввод чисел.")
		return
	}

	var rez float64
	switch oper {
	case "+":
		rez = a + b
	case "-":
		rez = a - b
	case "*":
		rez = a * b
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль")
			return
		}
		rez = a / b
	case "^":
		rez = math.Pow(a, b)
	default:
		fmt.Println("Недопустимая операция")
		return
	}
	fmt.Print(rez, "\n")
}

func zadacha2_2() {
	var (
		text, filteredtext string
	)
	fmt.Print("Введите текст:")
	text, _ = reader.ReadString('\n')
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filteredtext += string(unicode.ToLower(r))
		}
	}
	n := len(filteredtext)
	for i := 0; i < n/2; i++ {
		if filteredtext[i] != filteredtext[n-1-i] {
			fmt.Println("False")
			return
		}
	}

	fmt.Println("True")

}

func zadacha2_3() {
	var segments [3][2]int
	for i := 0; i < 3; i++ {
		fmt.Printf("Введите значение отрезка %d через пробел: ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Неверный формат ввода")
			return
		}
		segmentStart, err1 := strconv.Atoi(parts[0])
		segmentEnd, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Ошибка ввода чисел")
			return
		}
		segments[i][0] = segmentStart
		segments[i][1] = segmentEnd
	}

	for i := range segments {
		if segments[i][0] > segments[i][1] {
			segments[i][0], segments[i][1] = segments[i][1], segments[i][0]
		}
	}

	start := segments[0][0]
	end := segments[0][1]
	for _, segment := range segments {
		if segment[0] > start {
			start = segment[0]
		}
		if segment[1] < end {
			end = segment[1]
		}
	}

	fmt.Println(start <= end)
}

func zadacha2_4() {
	fmt.Println("Введите предложение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	textWithoutPunctuation := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, input)

	words := strings.Fields(textWithoutPunctuation)

	max := 0
	for _, word := range words {

		if len(word) >= max {
			max = len(word)
		}
	}

	for _, word := range words {
		if len(word) == max {
			fmt.Println(word)
			return
		}
	}
}

func zadacha2_5() {
	fmt.Print("Введите год:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	year, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("Год високосный")
		return
	}
	fmt.Println("Год невисокосный")

}

func zadacha3_1() {
	fmt.Print("Введите число n: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	a, b := 0, 1
	fmt.Print("Числа Фибоначчи: ")
	for a <= n {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimesInRange(start, end int) []int {
	var primes []int
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func zadacha3_2() {
	fmt.Print("Введите два числа через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Ошибка ввода")
		return
	}
	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка ввода чисел")
		return
	}

	primes := findPrimesInRange(start, end)
	fmt.Println("Простые числа в диапазоне:", primes)
}

func isArmstrong(num int) bool {
	numStr := strconv.Itoa(num)
	n := len(numStr)
	sum := 0
	for _, digit := range numStr {
		d, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(d), float64(n)))
	}
	return sum == num
}

func findArmstrongNumbersInRange(start, end int) []int {
	var armstrongNumbers []int
	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			armstrongNumbers = append(armstrongNumbers, i)
		}
	}
	return armstrongNumbers
}

func zadacha3_3() {
	fmt.Print("Введите два числа через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Ошибка ввода")
		return
	}
	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка ввода чисел")
		return
	}

	armstrongNumbers := findArmstrongNumbersInRange(start, end)
	fmt.Println("Числа Армстронга в диапазоне:", armstrongNumbers)
}

func zadacha3_4() {
	fmt.Println("Введите строку:")
	text, _ := reader.ReadString('\n')
	runes := []rune(text)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	fmt.Println(string(runes))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func zadacha3_5() {
	fmt.Print("Введите два числа через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) != 2 {
		fmt.Println("Ошибка ввода")
		return
	}
	a, err1 := strconv.Atoi(parts[0])
	b, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка ввода чисел")
		return
	}

	result := gcd(a, b)
	fmt.Printf("Наибольший общий делитель: %d\n", result)
}

func main() {
	for {
		fmt.Print("Введите номер задачи через тире (1-1,1-2), для выхода введите Q:")
		n, _ := reader.ReadString('\n')
		n = strings.TrimSpace(n)
		switch n {
		case "1-1":
			zadacha1_1()
		case "1-2":
			zadacha1_2()
		case "1-3":
			zadacha1_3()
		case "1-4":
			zadacha1_4()
		case "1-5":
			zadacha1_5()
		case "2-1":
			zadacha2_1()
		case "2-2":
			zadacha2_2()
		case "2-3":
			zadacha2_3()
		case "2-4":
			zadacha2_4()
		case "2-5":
			zadacha2_5()
		case "3-1":
			zadacha3_1()
		case "3-2":
			zadacha3_2()
		case "3-3":
			zadacha3_3()
		case "3-4":
			zadacha3_4()
		case "3-5":
			zadacha3_5()
		case "Q", "q", "Й", "й":
			os.Exit(0)
		default:
			fmt.Println("Введен неверный номер задачи")
		}
	}
}
