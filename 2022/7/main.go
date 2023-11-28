package main

import (
	"bufio"
	"fmt"
	"github.com/r4ndsen/adventofcode/support"
	"io"
	"log"
	"os"
	"strconv"
)

type NodeInterface interface {
	Size() int
	Name() string
}

type Node struct {
	root *Directory
	name string
}

type File struct {
	Node
	size int
}

type Directory struct {
	Node
	nodes []NodeInterface
}

func (n Node) Name() string {
	return n.name
}

type Result struct {
	sum int
}

func (d *Directory) Size100k(res *Result) {
	if d.Size() < 100000 {
		res.sum += d.Size()
		fmt.Printf("%q %d\n", d.name, d.Size())
	}

	for _, sub := range d.nodes {
		if sub, ok := sub.(*Directory); ok {
			sub.Size100k(res)
		}
	}
}

func (d *Directory) Size() int {
	sum := 0
	for _, sub := range d.nodes {
		sum += sub.Size()
	}

	return sum
}

func (f *File) Size() int {
	return f.size
}

func (d *Directory) addFile(size int, name string) {
	n := new(File)
	n.size = size
	n.name = name
	n.root = d

	d.nodes = append(d.nodes, n)
}

func (d *Directory) addDir(name string) {
	//fmt.Printf("add dir %q to %q\n", name, d.name)

	d.nodes = append(d.nodes, MkDir(name, d))
}

func (d *Directory) AddNode(line []byte) {
	if string(line[:3]) == "dir" {
		d.addDir(string(line[4:]))
		return
	}

	res := support.Split(' ', line)

	size, _ := strconv.Atoi(string(res[0]))

	d.addFile(size, string(res[1]))
}

func (n *Directory) chDir(name []byte) *Directory {
	//fmt.Println("chdir to " + string(name))

	if string(name) == ".." {
		//fmt.Printf("return to base dir %q\n", n.root.name)
		return n.root
	}

	if string(name) == "/" {
		if n.name == "/" {
			return n
		}

		return n.root.chDir(name)
	}

	for _, sub := range n.nodes {
		if sub.Name() == string(name) {
			return sub.(*Directory)
		}
	}

	log.Fatalf("child directory %q not found", string(name))

	return nil
}

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	currentNode := MkDir("/", nil)

	rootNode := currentNode

	for {
		line, err := r.ReadBytes(byte('\n'))

		if err == io.EOF {
			break
		}

		if len(line) == 0 {
			break
		}

		line = line[:len(line)-1]

		if line[0] != '$' {
			currentNode.AddNode(line)
		}

		if string(line[2:4]) == "ls" {
			continue
		}

		if string(line[2:4]) == "cd" {
			currentNode = currentNode.chDir(line[5:])
		}
	}

	fmt.Println(rootNode.Size())

	sum100k := new(Result)
	rootNode.Size100k(sum100k)
	fmt.Println(sum100k.sum)
}

func MkDir(name string, root *Directory) *Directory {
	d := new(Directory)
	d.root = root
	d.name = name
	d.nodes = make([]NodeInterface, 0)
	return d
}
