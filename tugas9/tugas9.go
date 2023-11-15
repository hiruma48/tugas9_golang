package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer pulihkan()
	var input string
	fmt.Println("masukan nomer:")
	fmt.Scanln(&input)

	var angka int
	var err error
	angka, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(angka, "ini adalah nomer")
	} else {
		fmt.Println(input, "bukan nomer")
		panic(err.Error())
		fmt.Println("saya muncul")
	}
}
func pulihkan() {
	if r := recover(); r != nil {
		fmt.Println("tidak bisa di tampilkan")
	} else {
		fmt.Println("sukses")
	}
}
