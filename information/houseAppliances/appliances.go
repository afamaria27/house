package houseAppliances

import "fmt"

type AppliancesStruct struct {
	Name  string
	Color string
}

//func AppliancesItems() []AppliancesStruct {
//	fridge := AppliancesStruct{"холодильник", "холодильник"}
//	oven := AppliancesStruct{"духовка", "холодильник"}
//	teapot := AppliancesStruct{"чайник", "холодильник"}
//	hairdryer := AppliancesStruct{"фен", "холодильник"}
//	blender := AppliancesStruct{"блендер", "холодильник"}
//	iron := AppliancesStruct{"утюг", "холодильник"}
//	return []AppliancesStruct{fridge, oven, teapot, hairdryer, blender, iron}
//}

func ShowAppliances(appliances []AppliancesStruct) {
	fmt.Println("\nБытовая техника:")
	for _, everyItem := range appliances {
		fmt.Print(everyItem.Name, " ", everyItem.Color, "\n")
	}
	fmt.Print("\n")
}
