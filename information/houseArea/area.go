package houseArea

import "fmt"

type AreaStruct struct {
	Name string
	Area float64
}

//func AreaItems() []AreaStruct {
//	bedroom := AreaStruct{"спальня", 25}
//	kitchen := AreaStruct{"кухня", 11.5}
//	bathroom := AreaStruct{"ванная комната", 10}
//	livingRoom := AreaStruct{"гостиная", 22}
//
//	return []AreaStruct{bedroom, kitchen, bathroom, livingRoom}
//}

func ShowArea(sizes []AreaStruct) {
	fmt.Print("\nПлощадь:\n")
	for _, size := range sizes {
		fmt.Print(size.Name, " ", size.Area, " кв.м", "\n")
	}
	fmt.Print("\n")
}
