package examples

import "fmt"

func UsingMaps() {
	fmt.Println("\nChapter 007 - Using Maps:")

	planetPos := make(map[string]int)
	planetPos["mercury"] = 1
	planetPos["venus"] = 2
	planetPos["earth"] = 3
	planetPos["mars"] = 4
	planetPos["jupiter"] = 5
	planetPos["saturn"] = 6
	planetPos["uranus"] = 7
	planetPos["neptune"] = 8
	fmt.Println("Position:", planetPos)

	planetTemp := map[string]int{
		"mercury": 270,
		"venus":   320,
		"Earth":   15,
		"Mars":    -25,
		"jupiter": -45,
		"saturn":  -68,
		"uranus":  -157,
		"neptune": -238,
	}
	fmt.Println("Temp:", planetTemp)

	//loop
	for i, v := range planetPos {
		fmt.Println("Planet:", i, "Position:", v)
	}

	//find a value
	value, exist := planetPos["pluto"]
	fmt.Println("pluto:", value, "exist:", exist)

	value, exist = planetPos["uranus"]
	fmt.Println("uranus:", value, "exist:", exist)

	fmt.Println("----------------------------------------------------------------")
}
