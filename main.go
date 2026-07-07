package main

import (
	"fmt"
)

type Depend struct {
	num int
	dep string // s - same, r - reverse
}

type Plate struct {
	pos    int
	depend []Depend
}

var one, seven int = 1, 7

func main() {
	var count int
	fmt.Print("Enter the count of plates: ")
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println("err_0:", err.Error())
		return
	}
	fmt.Scanln() // clear buffer

	plates := make([]Plate, count)
	fmt.Println("Enter the position of the pins on the plates, starting from the one closest to you")
	for i := 0; i < count; i++ {
		fmt.Printf("Enter the position of the %d pin: ", i+1)
		_, err := fmt.Scanf("%d", &plates[i].pos)
		if err != nil {
			fmt.Println("err_1:", err)
		}
		fmt.Scanln() // clear buffer

		fmt.Print("Enter the count of dependent plates: ")
		var dependence_count int
		_, err = fmt.Scanf("%d", &dependence_count)
		if err != nil {
			fmt.Println("err_2:", err)
		}
		fmt.Scanln() // clear buffer
		for j := 0; j < dependence_count; j++ {
			fmt.Print("Enter the number of dependent plate: ")
			var dep_plate int
			_, err = fmt.Scanf("%d", &dep_plate)
			if err != nil {
				fmt.Println("err_3:", err)
			}
			fmt.Scanln() // clear buffer

			fmt.Print("Enter the type of dependence (S/s - same, R/r - reverse): ")
			var type_of_dep string
			_, err = fmt.Scanf("%s", &type_of_dep)
			if err != nil {
				fmt.Println("err_4:", err)
			}
			fmt.Scanln() // clear buffer

			temp := Depend{num: dep_plate, dep: type_of_dep}
			plates[i].depend = append(plates[i].depend, temp)
		}
		fmt.Println()
	}

	// communications := make(map[int]int)

	fmt.Println("=============================")
	fmt.Println("CHECK THE WORK")
	fmt.Println("Count:", count)
	for i := 0; i < count; i++ {
		fmt.Printf("Position of the pin on the %d plate: %d\n", i+1, plates[i].pos)
		fmt.Println("Dependent plates and type of dependence:", plates[i].depend)
	}
	fmt.Println("=============================")

}
