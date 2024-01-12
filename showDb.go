package main

import (
	"house/information/houseAppliances"
	"house/information/houseArea"
	"house/information/houseFamily"
	"house/information/houseFurniture"
)

func ShowHouseDb() {
	frns := DbFurniture()
	fams := DbFamily()
	apps := DbAppliances()
	ars := DbArea()

	houseFurniture.ShowFurniture(frns)
	houseFamily.ShowFamily(fams)
	houseAppliances.ShowAppliances(apps)
	houseArea.ShowArea(ars)
}
