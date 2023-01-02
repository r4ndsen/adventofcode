package main

import (
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"log"
	"sort"
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
		if *v == *other {
			return 0
		} else if *v < *other {
			return 1
		} else {
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
		for i, v := range p.signals {
			if other.Len() <= i {
				return -1 // other does not have that many signals
			}
			if res := v.CompareTo(other.signals[i]); res != 0 {
				return res
			}
		}
		if p.Len() < other.Len() {
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

type PacketList []*Packet

func (p PacketList) Len() int {
	return len(p)
}

func (p PacketList) Less(i, j int) bool {
	return p[i].CompareTo(p[j]) == 1
}

func (p PacketList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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
	s := ValueSignal(value)
	p.signals = append(p.signals, &s)
}

func (p PacketList) String() string {
	var s strings.Builder

	for i, v := range p {
		if i != 0 {
			s.WriteString("\n")
		}
		s.WriteString(v.String())
	}

	return s.String()
}

func main() {
	key1, key2 := "[[2]]", "[[6]]"

	var packetList = PacketList{
		MakePacketStructure([]byte(key1)),
		MakePacketStructure([]byte(key2)),
	}

	for _, line := range support.GetInputFor(13) {
		if len(line) == 0 {
			continue
		}

		packetList = append(packetList, MakePacketStructure(line))
	}

	sort.Sort(packetList)

	decoderkey := 1

	for i, p := range packetList {
		if p.String() == key1 || p.String() == key2 {
			decoderkey *= i + 1
		}
	}
	fmt.Println("decoder key:", decoderkey)
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
