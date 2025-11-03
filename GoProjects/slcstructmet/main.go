package main

func main() {
	employees := PersonList{{Age: 18}, {Age: 44}, {Age: 18}}
	employees.GetAgePopularity()
}

type Person struct {
	Age uint8
}
type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	plm := make(map[uint8]int)
	for _, v := range pl {
		plm[v.Age]++
	}
	return plm

}
