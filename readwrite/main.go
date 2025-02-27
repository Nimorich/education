package main

import (
	"fmt"
	"io"
	"os"
)

type user struct {
	name string
	age  int
}

func (u *user) getUser() (n string, a int) {
	return u.name, u.age
}

type users interface {
	getUser()
}

func main() {
	u := user{"Vasya", 26}
	u.getUser()
	file := createFile()
	writeFile(&file)
	defer file.Close()
}
func readFile() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
	defer file.Close()
}
func writeFile(file *os.File) {
	date := ""
	name := ""
	fmt.Print("На какое число вы хотите забронировать номер?: ")
	fmt.Scan(&date)
	fmt.Print("Ваше имя: ")
	fmt.Scan(&name)
	file.WriteString(date + " " + name + "\n")
	fmt.Println("................................................")
	fmt.Printf("Поздравляю %s. Номер забронирован на: %s\n", name, date)
	file.Close()
}
func createFile() os.File {
	_, err := os.Stat("hello.txt")
	if err != nil {
		file, err := os.Create("hello.txt")
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		return *file
	} else {
		fileStat, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		return *fileStat
	}

}
