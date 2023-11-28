package main

import (
	"fmt"
)

type Monkey struct {
	id              int
	items           []int
	divisor         int
	operation       func(i int) int
	onTrue          *Monkey
	onFalse         *Monkey
	inspectionCount int
}

func (m *Monkey) decide(i int) bool {
	return i%m.divisor == 0
}

func (m *Monkey) throwItemToOtherMonkey(value int, other *Monkey) {
	fmt.Printf("    Item with worry level %v is thrown to monkey %v.\n", value, other.id)
	other.items = append(other.items, value)
}

func MostActiveMonkeys(monkeys []*Monkey) []*Monkey {
	var result []*Monkey

	var top1 *Monkey
	var top2 *Monkey

	for _, m := range monkeys {
		if top1 == nil {
			top1 = m
			continue
		}
		if top2 == nil {
			top2 = m
			continue
		}

		if m.inspectionCount > top1.inspectionCount {
			top2 = top1
			top1 = m
			continue
		}

		if m.inspectionCount > top2.inspectionCount {
			top2 = m
		}
	}

	return append(result, top1, top2)
}

func (m *Monkey) checkItems() {

	fmt.Printf("Monkey %v:\n", m.id)

	for len(m.items) > 0 {
		m.inspectionCount++
		item := m.items[:1][0]
		m.items = m.items[1:]

		fmt.Printf("  Monkey inspects item with worry level of %v.\n", item)
		val := m.operation(item)
		fmt.Printf("    Worry level increased from %v to %v.\n", item, val)

		decision := m.decide(val)
		fmt.Printf("    decision is %v.\n", decision)

		val /= 3
		fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %v.\n", val)

		if decision {
			m.throwItemToOtherMonkey(val, m.onTrue)
		} else {
			m.throwItemToOtherMonkey(val, m.onFalse)
		}

	}
}

func MakeMonkeys() []*Monkey {

	monkeys := make([]*Monkey, 8)

	for i := 0; i < 8; i++ {
		monkeys[i] = &Monkey{}
	}

	monkeys[0].items = []int{52, 60, 85, 69, 75, 75}
	monkeys[1].items = []int{96, 82, 61, 99, 82, 84, 85}
	monkeys[2].items = []int{95, 79}
	monkeys[3].items = []int{88, 50, 82, 65, 77}
	monkeys[4].items = []int{66, 90, 59, 90, 87, 63, 53, 88}
	monkeys[5].items = []int{92, 75, 62}
	monkeys[6].items = []int{94, 86, 76, 67}
	monkeys[7].items = []int{57}

	monkeys[0].divisor = 13
	monkeys[1].divisor = 7
	monkeys[2].divisor = 19
	monkeys[3].divisor = 2
	monkeys[4].divisor = 5
	monkeys[5].divisor = 3
	monkeys[6].divisor = 11
	monkeys[7].divisor = 17

	monkeys[1].id = 1
	monkeys[2].id = 2
	monkeys[3].id = 3
	monkeys[4].id = 4
	monkeys[5].id = 5
	monkeys[6].id = 6
	monkeys[7].id = 7

	monkeys[0].onTrue = monkeys[6]
	monkeys[1].onTrue = monkeys[0]
	monkeys[2].onTrue = monkeys[5]
	monkeys[3].onTrue = monkeys[4]
	monkeys[4].onTrue = monkeys[1]
	monkeys[5].onTrue = monkeys[3]
	monkeys[6].onTrue = monkeys[5]
	monkeys[7].onTrue = monkeys[6]

	monkeys[0].onFalse = monkeys[7]
	monkeys[1].onFalse = monkeys[7]
	monkeys[2].onFalse = monkeys[3]
	monkeys[3].onFalse = monkeys[1]
	monkeys[4].onFalse = monkeys[0]
	monkeys[5].onFalse = monkeys[4]
	monkeys[6].onFalse = monkeys[2]
	monkeys[7].onFalse = monkeys[2]

	monkeys[0].operation = func(i int) int { return i * 17 }
	monkeys[1].operation = func(i int) int { return i + 8 }
	monkeys[2].operation = func(i int) int { return i + 6 }
	monkeys[3].operation = func(i int) int { return i * 19 }
	monkeys[4].operation = func(i int) int { return i + 7 }
	monkeys[5].operation = func(i int) int { return i * i }
	monkeys[6].operation = func(i int) int { return i * 1 }
	monkeys[7].operation = func(i int) int { return i + 2 }

	return monkeys
}
