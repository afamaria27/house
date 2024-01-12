package houseFamily

import "fmt"

type FamilyStruct struct {
	Name string
	Role string
	Age  int
}

//func FamilyMembers() []FamilyStruct {
//	mother := FamilyStruct{"мама", "Лариса", 34}
//	father := FamilyStruct{"папа", "Олег", 71}
//	sister := FamilyStruct{"сестра", "Саша", 17}
//	dog := FamilyStruct{"собака", "Моцарт", 2}
//	turtle := FamilyStruct{"черепаха", "Череп", 10}
//	return []FamilyStruct{mother, father, sister, dog, turtle}
//}

func ShowFamily(family []FamilyStruct) {
	fmt.Println("\nСемья:")
	for _, person := range family {
		fmt.Print(person.Role, " ", person.Age, " ", person.Name, "\n")
	}
	fmt.Print("\n")
}
