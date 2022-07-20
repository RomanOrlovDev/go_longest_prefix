package main

import "fmt"

func main() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"парам", "пам", "папа"}
	prefix := ""

	counter := 0
outer:
	for _, character := range strs[0] {
		for _, str := range strs {
			str := []rune(str)
			if str[counter] != character {
				break outer
			}
		}
		prefix += string(character)
		counter++
	}
	fmt.Println(prefix)
}
