package houseAppliances

import "fmt"

type appliancesStruct struct {
	name string
}

func AppliancesItems() []appliancesStruct {
	fridge := appliancesStruct{"холодильник"}
	oven := appliancesStruct{"духовка"}
	teapot := appliancesStruct{"чайник"}
	hairdryer := appliancesStruct{"фен"}
	blender := appliancesStruct{"блендер"}
	iron := appliancesStruct{"утюг"}
	return []appliancesStruct{fridge, oven, teapot, hairdryer, blender, iron}
}

func ShowAppliances(appliances []appliancesStruct) {
	fmt.Println("\nБытовая техника:")
	for _, everyItem := range appliances {
		fmt.Print(everyItem.name, " ")
	}
	fmt.Print("\n")
}
