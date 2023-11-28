package main

import (
	"fmt"
	"strings"
)

type ValveCache map[string]*Valve

func (vc ValveCache) has(name string) bool {
	if _, ok := vc[name]; ok {
		return true
	}

	return false
}

func (vc ValveCache) addBlank(name string) {
	vc.add(NewValue(name))
}

func (vc ValveCache) get(name string) *Valve {
	if !vc.has(name) {
		vc.addBlank(name)
	}

	return vc[name]
}

func (vc ValveCache) add(v *Valve) {
	vc[v.name] = v
}

type Valve struct {
	valves   []*Valve
	name     string
	flowRate int
	open     bool
}

func (v *Valve) String() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Value %s has a flow rate=%d; tunnels lead to valves ", v.name, v.flowRate))

	for i, other := range v.valves {
		if i != 0 {
			s.WriteString(", ")
		}
		s.WriteString(other.name)
	}

	return s.String()
}

func (v *Valve) Add(other *Valve) {
	v.valves = append(v.valves, other)
}

func NewValue(name string) *Valve {
	return &Valve{
		name:   name,
		valves: make([]*Valve, 0),
	}
}
