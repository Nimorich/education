package main

import "fmt"

type Printer interface {
	Print()
}

// нигде не указано, что User реализует интерфейс Printer
type User struct {
	email string
}

// структура User имеет метод Print, как в интерфейсе Printer
// Следовательно, во время компиляции запишется связь между User и Print
func (u *User) Print() {
	fmt.Println("My email is", u.email)
}

// функция принимает как аргумент интерфейс Printer
func TestPrint(p Printer) {
	p.Print()
}

// в функцию TestPrint передается структура User,
// и так как она реализует интерфейс Printer, все работает без ошибок
func main() {
    fmt.Println("Hello")
    TestPrint(&User{email: "test@test.com"})
	//u := User{email: "testik@testail.com"}
	//DoHTTPCall(u.email)
}

type error interface {
	Error() string
}

type TimeoutErr struct {
	msg string
}

func (e *TimeoutErr) Error() string {
	return e.msg
}

// функция возвращает несколько аргументов, и ошибка возвращается последней
// func DoHTTPCall(e string) (string, error) {
// 	err := error
// 	if e == ""{
// 		return _, "nil"
// 	}
// 	return e, error
// }
