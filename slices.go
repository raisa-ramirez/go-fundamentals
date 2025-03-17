package main

import "fmt"

func main() {
	var groupA []string = make([]string, 3)
	groupA[0] = "Ana"
	groupA[1] = "Pao"
	groupA[2] = "Fa"

	fmt.Println(groupA)
	fmt.Println(groupA[2])

	groupB := make([]string, len(groupA))
	// copy groupA in groupB
	copy(groupB, groupA)
	fmt.Println(groupB)
	fmt.Println(groupB[0:2])
	fmt.Println(groupB[:1])
	fmt.Println(groupB[1:])

	languages := make([][]string, 3)
	languages[0] = make([]string, 2)
	languages[1] = make([]string, 3)

	languages[0][0] = "JAVA"
	languages[0][1] = "C#"
	languages[1][0] = "Go"
	languages[1][1] = "PHP"
	languages[1][2] = "JS"

	fmt.Println(languages)
}
