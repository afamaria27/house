package houseFamily

import "fmt"

type familyStruct struct {
	name string
}

func FamilyMembers() []familyStruct {
	mother := familyStruct{"мама"}
	father := familyStruct{"папа"}
	sister := familyStruct{"сестра"}
	dog := familyStruct{"собака"}
	turtle := familyStruct{"черепаха"}
	return []familyStruct{mother, father, sister, dog, turtle}
}

func ShowFamily(family []familyStruct) {
	fmt.Println("\nСемья:")
	for _, person := range family {
		fmt.Print(person.name, " ")
	}
}
