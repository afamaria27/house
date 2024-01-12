package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"house/information/houseAppliances"
	"house/information/houseArea"
	"house/information/houseFamily"
	"house/information/houseFurniture"
	"log"
)

func DbFamily() []houseFamily.FamilyStruct {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT role, age, name FROM family")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fams := make([]*houseFamily.FamilyStruct, 0)
	for rows.Next() {
		fam := new(houseFamily.FamilyStruct)
		err := rows.Scan(&fam.Role, &fam.Age, &fam.Name)
		if err != nil {
			log.Fatal(err)
		}
		fams = append(fams, fam)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	finalFamily := make([]houseFamily.FamilyStruct, len(fams))
	for i, fam := range fams {
		finalFamily[i] = houseFamily.FamilyStruct{
			Role: fam.Role,
			Age:  fam.Age,
			Name: fam.Name,
		}
	}
	return finalFamily
}

func DbFurniture() []houseFurniture.FurnitureStruct {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, material, color FROM furniture")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	frns := make([]*houseFurniture.FurnitureStruct, 0)
	for rows.Next() {
		frn := new(houseFurniture.FurnitureStruct)
		err := rows.Scan(&frn.Name, &frn.Material, &frn.Color)
		if err != nil {
			log.Fatal(err)
		}
		frns = append(frns, frn)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	finalFurniture := make([]houseFurniture.FurnitureStruct, len(frns))
	for i, frn := range frns {
		finalFurniture[i] = houseFurniture.FurnitureStruct{
			Name:     frn.Name,
			Material: frn.Material,
			Color:    frn.Color,
		}
	}
	return finalFurniture
}

func DbAppliances() []houseAppliances.AppliancesStruct {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color FROM appliances")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	apps := make([]*houseAppliances.AppliancesStruct, 0)
	for rows.Next() {
		app := new(houseAppliances.AppliancesStruct)
		err := rows.Scan(&app.Name, &app.Color)
		if err != nil {
			log.Fatal(err)
		}
		apps = append(apps, app)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	finalAppliances := make([]houseAppliances.AppliancesStruct, len(apps))
	for i, app := range apps {
		finalAppliances[i] = houseAppliances.AppliancesStruct{
			Name:  app.Name,
			Color: app.Color,
		}
	}
	return finalAppliances
}

func DbArea() []houseArea.AreaStruct {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, area FROM sizes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ars := make([]*houseArea.AreaStruct, 0)
	for rows.Next() {
		size := new(houseArea.AreaStruct)
		var err = rows.Scan(&size.Name, &size.Area)
		if err != nil {
			log.Fatal(err)
		}
		ars = append(ars, size)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	finalArea := make([]houseArea.AreaStruct, len(ars))
	for i, size := range ars {
		finalArea[i] = houseArea.AreaStruct{
			Name: size.Name,
			Area: size.Area,
		}
	}
	return finalArea
}
