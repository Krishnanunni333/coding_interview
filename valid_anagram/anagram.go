package main

import "fmt"

func checkAnagram(str1 string, str2 string) bool{
	if len(str1) != len(str2){
		return false
	}
	freq := make(map[rune]int)
	for _, char := range str1{
		freq[char]++
	}

	for _, char := range str2{
		freq[char]--
	}
	for _, value := range freq{
		if value != 0{
			return false
		}
	}

	return true
}
func main(){
	str1 := "d"
	str2 := ""
	fmt.Println(checkAnagram(str1, str2))
}