package main
import "fmt"
func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		{
			return 0
		}
	case "CANCELLED":
		{
			return 1
		}
	case "UNKNOWN":
		{
			return 2
		}
	default:
		{
			return 2
		}
	}
}

func main(){
	fmt.Println("Hello emacs")

}
