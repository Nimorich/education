package main

func main() {

}
func Map(strs []string, mapFunc func(s string) string) []string {
	var rslice []string = make([]string, len(strs))
	for i := 0; i < len(strs); i++ {
		rslice[i] = mapFunc(strs[i])
	}
	return rslice
}
