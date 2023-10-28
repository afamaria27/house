package createShowHouse

import (
	"house/information/houseAppliances"
	"house/information/houseArea"
	"house/information/houseFamily"
	"house/information/houseFurniture"
)

func ShowHouse() {
	area := houseArea.AreaItems()
	furniture := houseFurniture.FurnitureItems()
	appliances := houseAppliances.AppliancesItems()
	family := houseFamily.FamilyMembers()

	houseArea.ShowArea(area)
	houseFurniture.ShowFurniture(furniture)
	houseAppliances.ShowAppliances(appliances)
	houseFamily.ShowFamily(family)
}
