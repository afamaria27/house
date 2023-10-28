package houseArea

import "fmt"

type areaStruct struct {
	name string
	area float64
}

func AreaItems() []areaStruct {
	bedroom := areaStruct{"спальня", 25}
	kitchen := areaStruct{"кухня", 11.5}
	bathroom := areaStruct{"ванная комната", 10}
	livingRoom := areaStruct{"гостиная", 10}

	return []areaStruct{bedroom, kitchen, bathroom, livingRoom}
}

func ShowArea(sizes []areaStruct) {
	fmt.Print("Площадь:\n")
	for _, size := range sizes {
		fmt.Print(size.name, " ", size.area, " кв.м", "\n")
	}
}
