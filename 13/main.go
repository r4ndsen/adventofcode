package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"log"
	"strings"
)

type Signal interface {
	String() string
	CompareTo(signal Signal) int
}

type ValueSignal int

func (p *Packet) Len() int {
	return len(p.signals)
}

func (v *ValueSignal) String() string {
	return fmt.Sprintf("%v", int(*v))
}

func (v *ValueSignal) CompareTo(other Signal) int {
	switch other := other.(type) {
	case *ValueSignal:
		fmt.Printf("Compare %v vs %v\n", v, other)
		if *v == *other {
			return 0
		} else if *v < *other {
			fmt.Println("Left side is smaller, so inputs are in the right order")
			return 1
		} else {
			fmt.Println("Right side is smaller, so inputs are not in the right order")
			return -1
		}
	case *Packet:
		asPackage := ValueSignalToPacket(v)
		return asPackage.CompareTo(other)
	default:
		log.Fatal("not a valid type", other)
	}

	return 0
}

func ValueSignalToPacket(v *ValueSignal) *Packet {
	p := NewPacket(nil)
	p.AddValueSignal(int(*v))

	return p
}

/**
 *  1 = correct
 *  0 = undecided
 * -1 = incorrect
 */
func (p *Packet) CompareTo(other Signal) int {
	switch other := other.(type) {
	case *Packet:
		fmt.Println("Compare", p, "to", other)
		for i, v := range p.signals {
			if other.Len() <= i {
				fmt.Println("Right side ran out of items, so inputs are not in the right order")
				return -1 // other does not have that many signals
			}
			if res := v.CompareTo(other.signals[i]); res != 0 {
				return res
			}
		}
		if p.Len() < other.Len() {
			fmt.Println("Left side ran out of items, so inputs are in the right order")
			return 1
		}
	case *ValueSignal:
		return p.CompareTo(ValueSignalToPacket(other))
	default:
		log.Fatal("not a valid type", other)
	}

	return 0
}

type Packet struct {
	signals []Signal
	parent  *Packet
}

func (p *Packet) String() string {
	var s strings.Builder
	s.WriteString("[")

	for i, v := range p.signals {
		if i != 0 {
			s.WriteString(",")
		}

		s.WriteString(v.String())
	}

	s.WriteString("]")

	return s.String()
}

func NewPacket(parent *Packet) *Packet {
	return &Packet{
		signals: make([]Signal, 0),
		parent:  parent,
	}
}

func (p *Packet) AddPacketSignal() *Packet {
	child := NewPacket(p)
	p.signals = append(p.signals, child)

	return child
}

func (p *Packet) AddValueSignal(value int) {
	signal := ValueSignal(value)
	p.signals = append(p.signals, &signal)
}

func main() {
	var right, left *Packet

	for _, line := range support.GetInputFor(13) {
		if len(line) == 0 {
			left = nil
			right = nil
			continue
		}

		if left == nil {
			left = MakePacketStructure(line)
			continue
		}

		right = MakePacketStructure(line)

		ComparePackets(left, right)
	}

	fmt.Println("comparesum:", compareSum)
}

var compareIndex int
var compareSum int

func ComparePackets(left, right *Packet) {
	compareIndex++

	res := left.CompareTo(right)
	switch res {
	case 1, 0:
		compareSum += compareIndex
	case -1:
	}
}

func MakePacketStructure(line []byte) *Packet {
	line = line[1 : len(line)-1] // cut off starting and ending "[" "]"
	currentPackage := NewPacket(nil)
	root := currentPackage

	var intBuf []byte

	for _, chr := range line {
		if support.IsInt(string(chr)) {
			intBuf = append(intBuf, chr)
			continue
		}

		if len(intBuf) != 0 {
			currentPackage.AddValueSignal(support.ToInt(string(intBuf)))
			intBuf = nil
		}

		if chr == '[' {
			currentPackage = currentPackage.AddPacketSignal()
			continue
		}
		if chr == ']' {
			currentPackage = currentPackage.parent
			continue
		}
	}

	if len(intBuf) != 0 {
		currentPackage.AddValueSignal(support.ToInt(string(intBuf)))
	}

	return root
}
