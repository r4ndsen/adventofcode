package main

import (
	"fmt"
)

func main() {

	monkeys := MakeMonkeys()

	for round := 1; round <= 20; round++ {
		fmt.Println("round", round)
		for _, m := range monkeys {
			m.checkItems()
			//		fmt.Printf("Monkey %v inspected items %v times.\n", m.id, m.inspectionCount)
		}
	}

	for _, m := range monkeys {
		fmt.Printf("Monkey %v inspected items %v times.\n", m.id, m.inspectionCount)
	}

	/*for round := 0; round < 19; round++ {
		fmt.Println("round", round+2)
		for _, m := range MostActiveMonkeys(monkeys) {
			m.checkItems()
		}
	}*/
}
