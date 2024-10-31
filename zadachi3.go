package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func triangleArea(base float64, height float64) float64 {
	return base * height / 2
}

func sortArray(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		if arr[i+1] < arr[i] {
			arr[i+1], arr[i] = arr[i], arr[i+1]
			swapped = true
		}
		if !swapped {
			break
		}
	}
	return arr
}

func sumOfSquares(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i * i
		}
	}
	return sum
}

func isPalidrome(s string) bool {
	stat := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			stat = false
		}
	}
	return stat
}

func isPrime(n int) bool {
	var prime []int
	for i := 2; i <= n; i++ {
		prime = append(prime, i)
	}
	for i := 2; i <= n; i++ {
		if prime[i-2] != 0 {
			for j := i + 1; j <= n; j++ {
				if j%i == 0 {
					prime[j-2] = 0
				}
			}
		}
	}
	return prime[len(prime)-1] == n
}

func generatePrimes(limit int) []int {
	var numbers, prime []int
	for i := 2; i <= limit; i++ {
		numbers = append(numbers, i)
	}
	for i := 2; i <= limit; i++ {
		if numbers[i-2] != 0 {
			for j := i + 1; j <= limit; j++ {
				if j%i == 0 {
					numbers[j-2] = 0
				}
			}
		}
	}
	for _, num := range numbers {
		if num != 0 {
			prime = append(prime, num)
		}
	}
	return prime
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func arrayToString(arr []int) string {
	var strArr []string
	for _, num := range arr {
		strArr = append(strArr, fmt.Sprintf("%d", num))
	}
	return strings.Join(strArr, "")
}

func toBinary(n int) string {
	var bin []int
	for n != 0 {
		bin = append(bin, n%2)
		n = n / 2
	}
	reverseBin := reverseArray(bin)

	return arrayToString(reverseBin)
}

func findMax(arr []int) int {
	max := arr[0]
	for _, r := range arr {
		if r > max {
			max = r
		}
	}
	return max
}

func gcd(a int, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func sumArray(arr []int) int {
	var sum int
	for _, r := range arr {
		sum += r
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите номер задачи: ")
	scanner.Scan()
	input := scanner.Text()
	switch input {
	case "1":
		fmt.Print("Введите высоту: ")
		scanner.Scan()
		height, _ := strconv.ParseFloat(scanner.Text(), 64)
		fmt.Print("Введите основание: ")
		scanner.Scan()
		base, _ := strconv.ParseFloat(scanner.Text(), 64)
		fmt.Print(triangleArea(base, height))
	case "2":
		fmt.Print("Введите числа массива через пробел: ")
		scanner.Scan()
		input := scanner.Text()
		text := strings.Fields(input)
		arr := []int{}
		for _, i := range text {
			element, _ := strconv.Atoi(i)
			arr = append(arr, element)
		}
		fmt.Print(sortArray(arr))
	case "3":
		fmt.Print("Введите целое число: ")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(sumOfSquares(n))
	case "4":
		fmt.Print("Введите строку: ")
		scanner.Scan()
		fmt.Println(isPalidrome(scanner.Text()))
	case "5":
		fmt.Print("Введите число: ")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(isPrime(n))
	case "6":
		fmt.Print("Введите число: ")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(generatePrimes(n))
	case "7":
		fmt.Print("Введите число: ")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(toBinary(n))
	case "8":
		fmt.Print("Введите числа массива через пробел: ")
		scanner.Scan()
		input := scanner.Text()
		text := strings.Fields(input)
		arr := []int{}
		for _, i := range text {
			element, _ := strconv.Atoi(i)
			arr = append(arr, element)
		}
		fmt.Println(findMax(arr))
	case "9":
		fmt.Print("Введите два числа через пробел: ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		fmt.Println(gcd(a, b))
	case "10":
		fmt.Print("Введите числа массива через пробел: ")
		scanner.Scan()
		input := scanner.Text()
		text := strings.Fields(input)
		arr := []int{}
		for _, i := range text {
			element, _ := strconv.Atoi(i)
			arr = append(arr, element)
		}
		fmt.Println(sumArray(arr))
	}

}
