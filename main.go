package main

import (
	"fmt"
	"strings"
)

func base36ToChar(number int) string {
	if number < 10 {
		return fmt.Sprint(number)
	}
	return string(rune(65 + number - 10))
}

func getPrintableHash2(values ...string) string {
	val1 := 0
	for _, value := range values {
		if len(value) > val1 {
			val1 = len(value)
		}
	}

	var result strings.Builder
	num1 := val1
	num2 := 1
	for num2 <= num1 {
		num3 := 0
		for _, value := range values {
			stringVal := string(value[(num2-1)%len(value)])
			num3 += int(stringVal[0])
		}

		number := num3 % 36
		result.WriteString(base36ToChar(number))
		num2++
	}

	return result.String()
}

func getPrintableHash(data string) string {
	var result strings.Builder
	for _, char := range data {
		number := int(char) % 36
		result.WriteString(base36ToChar(number))
	}
	return result.String()
}

func secretHash(activationID, serialNo string) string {
	activationID = strings.ToUpper(activationID)
	serialNo = strings.ToUpper(serialNo)

	if len(activationID) != 9 {
		return getPrintableHash(activationID)
	}

	if activationID[0] == '1' {
		return getPrintableHash2(activationID, serialNo, "PaloAlto.com")
	}

	panic("Invalid Activation Key")
}

func isActivationCodeValid(serialNumber, activationKey string) string {
	return secretHash(activationKey, serialNumber)[:8]
}

func main() {
	fmt.Println("Welcome to  Bussines Plan Pro and Sales and Marketing Pro activator ")
	fmt.Println("Please enter your serial number:")
	var serialNumber string
	fmt.Scanln(&serialNumber)
	serialNumber = strings.ReplaceAll(serialNumber, "-", "")
	fmt.Println("Please enter your activation ID:")
	var activationCode string
	fmt.Scanln(&activationCode)
	fmt.Println("Your activation key is:")
	fmt.Println(isActivationCodeValid(serialNumber, activationCode))
	fmt.Println("Press enter to exit")
	fmt.Scanln()
}
