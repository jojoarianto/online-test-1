package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {

		// get m (size array vertical)
		var n int
		fmt.Scanln(&n)

		// get n (size array horizantal)
		var m int
		fmt.Scanln(&m)

		// create array
		words := make([][]string, n)
		for i := range words {
			words[i] = make([]string, m)
		}

		var army []string
		var armyLocX []int
		var armyLocY []int
		// var armyQty []int

		/**
		 * I use 2d array to represent the map
		 *
		 * description :
		 *
		 * maps[row][col]
		 *
		 * n = row
		 * m = col
		 *
		 * value: 0 = not visited node
		 * 		  1 = mountain
		 * 		  2 = visited node
		 *
		 * Stack status :
		 */
		for i := 0; i < n; i++ {
			var str string
			fmt.Scanln(&str)

			for idx, r := range str {
				dt := string(r)
				words[i][idx] = dt

				// find army
				if "#" != dt && "." != dt {
					army = append(army, dt)
					armyLocX = append(armyLocX, idx)
					armyLocY = append(armyLocY, i)
				}
			}
		}

		fmt.Println(army)
		fmt.Println(armyLocX)
		fmt.Println(armyLocY)

		// find army

	}

	/**
	 * using depth search algorithm
	 *
	 * move to down
	 * move to right
	 * move to up
	 * move to left
	 */
	func travelFindAnotherArmy(army string, armyLocX int, armyLocY) {
		
		
	}

}
