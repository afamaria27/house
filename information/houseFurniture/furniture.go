package houseFurniture

import "fmt"

type FurnitureStruct struct {
	Name     string
	Material string
	Color    string
}

//func FurnitureItems() []FurnitureStruct {
//	bed := FurnitureStruct{"кровать", "дерево", "коралловый"}
//	couch := FurnitureStruct{"диван", "кожа", "фиолетовый"}
//	closet := FurnitureStruct{"шкаф", "фанера", "желтый"}
//	table := FurnitureStruct{"стол", "стекло", "белый"}
//	armchair := FurnitureStruct{"кресло", "ткань", "бежевый"}
//	return []FurnitureStruct{bed, couch, closet, table, armchair}
//}

func ShowFurniture(furniture []FurnitureStruct) {
	fmt.Println("\nМебель:")
	for _, everyItem := range furniture {
		fmt.Print(everyItem.Name, " ", everyItem.Material, " ", everyItem.Color, "\n")
	}
	fmt.Print("\n")
}
