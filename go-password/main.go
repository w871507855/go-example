package main

import "fmt"

func main() {
	password := "123456"
	passwordMd5 := GeneratePassword(password)
	fmt.Println("加密后的密码", passwordMd5)
	result := ComparePassword(passwordMd5, password)
	fmt.Println("密码是否相同", result)
}
