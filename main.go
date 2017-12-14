package main

import (
	"fmt"
	"math"
)

type puzzleSolution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolveZebraPuzzle() puzzleSolution {
	nationalities := []string{"english", "spaniard", "norwegian", "ukrainian", "japanese"}
	colors := []string{"red", "green", "yellow", "blue", "ivory"}
	pets := []string{"dog", "horse", "fox", "snails", "zebra"}
	drinks := []string{"tea", "coffee", "juice", "milk", "water"}
	cigarretes := []string{"gold", "kools", "chester", "lucky", "parliaments"}

	var chosenNats = []string{"", "", "", "", ""}
	var chosenColors = []string{"", "", "", "", ""}
	var chosenPets = []string{"", "", "", "", ""}
	var chosenDrinks = []string{"", "", "", "", ""}
	var chosenCigs = []string{"", "", "", "", ""}

	for _, nationalityList := range permutations(nationalities) {
		// 10. The Norwegian lives in the first house.
		if nationalityList[0] != "norwegian" {
			continue
		}

		for _, colorList := range permutations(colors) {
			// 15. The Norwegian lives next to the blue house.
			if colorList[1] != "blue" {
				continue
			}

			// 2. There are five houses
			if indexOf("english", nationalityList) != indexOf("red", colorList) {
				continue
			}

			// 6. The green house is immediately to the right of the ivory house.
			if indexOf("ivory", colorList)+1 != indexOf("green", colorList) {
				continue
			}

			for _, drinkList := range permutations(drinks) {
				// 9. Milk is drunk in the middle house.
				if drinkList[2] != "milk" {
					continue
				}

				// 4. Coffee is drunk in the green house.
				if indexOf("coffee", drinkList) != indexOf("green", colorList) {
					continue
				}

				// 5. The Ukrainian drinks tea.
				if indexOf("tea", drinkList) != indexOf("ukrainian", nationalityList) {
					continue
				}

				for _, cigList := range permutations(cigarretes) {
					// 8. Kools are smoked in the yellow house.
					if indexOf("kools", cigList) != indexOf("yellow", colorList) {
						continue
					}
					// 13. The Lucky Strike smoker drinks orange juice.
					if indexOf("lucky", cigList) != indexOf("juice", drinkList) {
						continue
					}
					// 14. The Japanese smokes Parliaments.
					if indexOf("parliaments", cigList) != indexOf("japanese", nationalityList) {
						continue
					}

					for _, petList := range permutations(pets) {
						// 3. The Spaniard owns the dog
						if indexOf("spaniard", nationalityList) != indexOf("dog", petList) {
							continue
						}
						// 7. The Old Gold smoker owns snails.
						if indexOf("gold", cigList) != indexOf("snails", petList) {
							continue
						}
						// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
						if math.Abs(float64(indexOf("chester", cigList)-indexOf("fox", petList))) != 1 {
							continue
						}
						// 12. Kools are smoked in the house next to the house where the horse is kept.
						if math.Abs(float64(indexOf("kools", cigList)-indexOf("horse", petList))) != 1 {
							continue
						}

						copy(chosenPets, petList)
						break
					}
					if len(chosenPets[0]) == 0 {
						continue
					}
					copy(chosenCigs, cigList)
					break
				}
				if len(chosenPets[0]) == 0 {
					continue
				}
				copy(chosenDrinks, drinkList)
				break
			}
			if len(chosenDrinks[0]) == 0 {
				continue
			}
			copy(chosenColors, colorList)
			break
		}
		if len(chosenColors[0]) == 0 {
			continue
		}
		copy(chosenNats, nationalityList)
		break
	}

	waterPos := indexOf("water", chosenDrinks)
	zebraPos := indexOf("zebra", chosenPets)

	return puzzleSolution{
		DrinksWater: chosenNats[waterPos],
		OwnsZebra:   chosenNats[zebraPos],
	}
}

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	solution := SolveZebraPuzzle()
	fmt.Println(solution.DrinksWater + " drinks water")
	fmt.Println(solution.OwnsZebra + " owns a zebra")
}
