package houseFurniture

import "fmt"

type furnitureStruct struct {
	name string
}

func FurnitureItems() []furnitureStruct {
	bed := furnitureStruct{"кровать"}
	couch := furnitureStruct{"диван"}
	closet := furnitureStruct{"шкаф"}
	table := furnitureStruct{"стол"}
	armchair := furnitureStruct{"кресло"}
	return []furnitureStruct{bed, couch, closet, table, armchair}
}

func ShowFurniture(furniture []furnitureStruct) {
	fmt.Println("\nМебель:")
	for _, everyItem := range furniture {
		fmt.Print(everyItem.name, " ")
	}
	fmt.Print("\n")
}
