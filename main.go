package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	//"strconv"
)

// File struct of files
type File struct {
	path string
	name string
	size string
}

// Node struct realizes node of graph
type Node struct {
	Files []File
	NodeName string
}

// Pair c++ style)))0)
type Pair struct {
	First *Node
	Second *Node
}

// ArcGraph struct realizes graph (in this realization Node = dir, dir contains list of files)
type ArcGraph struct {
	Nodes []*Pair
	Size int64
}

// getNextDirs returns mas pairs next verts from vertex
func (s *ArcGraph) getNextDirs(vertex string) []Pair {
	nextDirs := make([]Pair, 5, 10)
	for _, node := range s.Nodes {
		if node.First.NodeName == vertex {
			nextDirs = append(nextDirs, *node)
		}
	}

	return nextDirs
}

// getNextDirs returns parent of vertex
func (s *ArcGraph) getPrevDir(vertex string) *Node {
	for _, node := range s.Nodes {
		if node.Second.NodeName == vertex {
			return node.First
		}
	}

	return nil
}

// AddVertex added a node to graph
func (s *ArcGraph) AddVertex(first *Node, second *Node) {
	s.Nodes = append(s.Nodes, &Pair{First:first, Second:second})
	s.Size++
}

// FindNode finds node by name and returns it ptr
func (s *ArcGraph) FindNode(name string) *Node {
	for _, pair := range s.Nodes {
		if pair.Second.NodeName == name {
			return pair.Second
		} else if pair.First.NodeName == name {
			return pair.First
		}
	}

	return nil
}


//////////////////////////////////////////
// https://flaviocopes.com/golang-data-structure-stack/

// StringStack stack of string
type StringStack struct {
	items []string
	// mutex
}

// New creates a new obj of StringStack
func (s *StringStack) New() *StringStack {
	s.items = []string{}
	return s
}

// Push adds item on the top of StringStack
func (s *StringStack) Push(t string) {
	s.items = append(s.items, t)
}

// Pop delete item from top and returning it
func (s *StringStack) Pop() string {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

// Len returns size of StringStack
func (s *StringStack) Len() int {
	return len(s.items)
}

///////////////////////////////////////////////


func dirTree(out *bytes.Buffer, path string, printFiles bool) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}


	for _, file := range files {
		if !file.IsDir() {
			out.WriteString("├───" + file.Name() + " (" + strconv.FormatInt(file.Size(), 10) + "b) \n")
		} else if file.IsDir() {
			out.WriteString("├───" + file.Name() + " ---new_dir\n")
			err = dirTree(out, path + "/" + file.Name(), printFiles)
			if err != nil {
				log.Fatal(err)
			}
		}

	}


	return nil
}

func main() {
	//out := os.Stdout
	//if !(len(os.Args) == 2 || len(os.Args) == 3) {
	//	panic("usage go run main.go . [-f]")
	//}
	//path := os.Args[1]
	//printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	//err := dirTree(out, path, printFiles)
	//if err != nil {
	//	panic(err.Error())
	//}
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", true)
	if err != nil {
		fmt.Println("test for OK Failed - error")
	}
	fmt.Println(out)
}
