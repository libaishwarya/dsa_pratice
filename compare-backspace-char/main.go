package main

import "fmt"

func main() {
	// s := "ab#c"
	// t := "ad#c"
	// fmt.Println(backspacecompare(s, t))

	s1 := "ab##"
	t1 := "c#d#"
	fmt.Println(backspacecompare(s1, t1))

	// s2 := "a#c"
	// t2 := "b"
	// fmt.Println(backspacecompare(s2, t2))

	// s3 := "a##c"
	// t3 := "#a#c"
	// fmt.Println(backspacecompare(s3, t3))
}

func backspacecompare(s, t string) bool {
	compare1 := []rune{}
	compare2 := []rune{}
	for _, value := range s {
		if value == '#' {
			if len(compare1) > 0 {
				compare1 = compare1[:len(compare1)-1] // Remove last character
			}
		} else {
			compare1 = append(compare1, value)
		}
	}

	for _, value := range t {
		if value == '#' {
			if len(compare2) > 0 {
				compare2 = compare2[:len(compare2)-1] // Remove last character
			}
		} else {
			compare2 = append(compare2, value)
		}
	}

	return string(compare1) == string(compare2)

}
