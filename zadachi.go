package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func sumOfDigits(input string) int {
	sum := 0
	for _, r := range input {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			//fmt.Println(err)
			return -1
		}
		sum += digit
	}
	return sum
}

func zadacha1_1() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")

	scanner.Scan()
	input := scanner.Text()

	result := sumOfDigits(input)
	fmt.Printf("Сумма цифр: '%d'\n", result)
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func zadacha1_2() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите температуру, не забудтье указать единицу иземерения: ")
	scanner.Scan()
	input := scanner.Text()

	if len(input) < 2 {
		fmt.Println("Некорректный ввод")
		return
	}

	value, err := strconv.ParseFloat(input[:len(input)-1], 64)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}

	unit := input[len(input)-1]
	switch unit {
	case 'C', 'c':
		fmt.Printf("%.2fC = %.2fF\n", value, celsiusToFahrenheit(value))
	case 'F', 'f':
		fmt.Printf("%.2fF = %.2fC\n", value, fahrenheitToCelsius(value))
	default:
		fmt.Println("Неизвестная единица измерения")
	}
}

func zadacha1_3() {
	nums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите цифры:")
	scanner.Scan()
	text := scanner.Text()
	for _, r := range text {
		digit, err := strconv.Atoi(string(r))
		if err == nil {
			nums = append(nums, digit)
		}
	}
	fmt.Printf("%v\n", nums)
	for i, r := range nums {
		nums[i] = r * 2
	}
	fmt.Printf("%v\n", nums)
}

func zadacha1_4() {
	text := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите первую строчку:")
	scanner.Scan()
	text1 := scanner.Text()
	text = append(text, text1)
	fmt.Printf("Введите вторую строчку:")
	scanner.Scan()
	text2 := scanner.Text()
	text = append(text, text2)
	fmt.Printf(strings.Join(text, " "))
	fmt.Printf("\n")
}

func zadacha1_5() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите координаты первой точки через пробел (например 1 2):")
	scanner.Scan()
	xt := scanner.Text()
	fmt.Printf("Введите координаты второй точки через пробел (например 1 2):")
	scanner.Scan()
	yt := scanner.Text()
	xt1 := string(xt[0])
	xt2 := string(xt[2])
	yt1 := string(yt[0])
	yt2 := string(yt[2])
	x1, _ := strconv.ParseFloat(xt1, 64)
	x2, _ := strconv.ParseFloat(xt2, 64)
	y1, _ := strconv.ParseFloat(yt1, 64)
	y2, _ := strconv.ParseFloat(yt2, 64)
	l := math.Sqrt(math.Pow(y1-x1, 2) + math.Pow(y2-x2, 2))
	fmt.Printf("Расстояние между точками: %.2f\n", l)
}

func zadacha2_1() {
	fmt.Printf("Введите число:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	digit, _ := strconv.Atoi(string(text))
	switch digit % 2 {
	case 1, -1:
		fmt.Printf("Число нечетное\n")
	case 0:
		fmt.Printf("Число четное\n")
	}

}

func zadacha2_2() {
	fmt.Print("Введите год:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	year, _ := strconv.Atoi(string(text))
	var stat bool = false
	if year%400 == 0 {
		stat = true
	}
	if year%4 == 0 {
		stat = true
	}
	if stat {
		fmt.Print("Год високосный\n")
	} else {
		fmt.Print("Год невисокосный\n")
	}

}

func zadacha2_3() {
	nums := []int{}
	numst := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите цифры через пробел:")
	scanner.Scan()
	text := scanner.Text()
	numst = append(numst, strings.Split(text, " ")...)
	for _, r := range numst {
		digit, err := strconv.Atoi(string(r))
		if err == nil {
			nums = append(nums, digit)
		}
	}
	max := nums[0]
	for _, r := range nums {
		if r > max {
			max = r
		}
	}
	fmt.Printf("Максимальное число: %d\n", max)
}

func zadacha2_4() {
	fmt.Print("Введите возраст:")
	var stat string = ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	age, _ := strconv.Atoi(text)
	if age < 12 && age >= 0 {
		stat = "Ребенок\n"
	} // ребенок от 0 до 11
	if age > 11 && age < 17 {
		stat = "Подросток\n"
	} // подросток от 12 до 16
	if age > 16 && age < 60 {
		stat = "Взрослый\n"
	} // взрослый от 17 до 60
	if age >= 60 {
		stat = "Пожилой\n"
	} // пожилой если старше 60
	fmt.Print(stat)
}

func zadacha2_5() {
	fmt.Print("Введите число:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	digit, _ := strconv.Atoi(text)
	if digit%3 == 0 && digit%5 == 0 {
		fmt.Print("Число делится на 3 и 5\n")
	} else {
		fmt.Print("Не делится\n")
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func zadacha3_1() {
	fmt.Print("Введите число: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	num, _ := strconv.Atoi(text)

	result := factorial(num)
	fmt.Printf("Факториал числа %d равен %d\n", num, result)
}

func fibonacci(n int) []int {
	fib := make([]int, n)
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func zadacha3_2() {
	fmt.Print("Введите количество чисел Фибоначчи: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	n, _ := strconv.Atoi(text)

	fibNumbers := fibonacci(n)
	fmt.Printf("Первые %d чисел Фибоначчи: %v\n", n, fibNumbers)
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func zadacha3_3() {
	nums := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите цифры:")
	scanner.Scan()
	text := scanner.Text()
	for _, r := range text {
		digit, err := strconv.Atoi(string(r))
		if err == nil {
			nums = append(nums, digit)
		}
	}
	reversed := reverseArray(nums)
	fmt.Printf("Реверс массива: %v\n", reversed)
}

func zadacha3_4() {
	prime := []int64{}
	fmt.Print("Введите число:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	n, _ := strconv.Atoi(text)
	var i int64
	for i = 1; i <= int64(n); i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			prime = append(prime, i)
		}
	}
	fmt.Printf("%v\n", prime)
}

func zadacha3_5() {
	nums := []int{}
	numst := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите чмсла через пробел:")
	scanner.Scan()
	text := scanner.Text()
	numst = append(numst, strings.Split(text, " ")...)
	for _, r := range numst {
		digit, err := strconv.Atoi(string(r))
		if err == nil {
			nums = append(nums, digit)
		}
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Printf("Сумма чисел %d\n", sum)
}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Введите номер задачи и номер раздела через пробел (например, 1 2), для выхода введите 'Q':")
		scanner.Scan()
		number := scanner.Text()
		switch number {
		case "1 1":
			zadacha1_1()
		case "1 2":
			zadacha1_2()
		case "1 3":
			zadacha1_3()
		case "1 4":
			zadacha1_4()
		case "1 5":
			zadacha1_5()
		case "2 1":
			zadacha2_1()
		case "2 2":
			zadacha2_2()
		case "2 3":
			zadacha2_3()
		case "2 4":
			zadacha2_4()
		case "2 5":
			zadacha2_5()
		case "3 1":
			zadacha3_1()
		case "3 2":
			zadacha3_2()
		case "3 3":
			zadacha3_3()
		case "3 4":
			zadacha3_4()
		case "3 5":
			zadacha3_5()
		case "Q", "q", "Й", "й":
			os.Exit(0)
		default:
			fmt.Printf("Задачи с таким номером нет\n")
		}
	}
}
