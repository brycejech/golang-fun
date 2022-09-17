package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type DirTree struct {
	Name  string
	Path  string
	IsDir bool
	Depth int
	Files []DirTree
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must provide path arg")
	}
	rootPath := os.Args[1]

	tree, _ := BuildDirTree(rootPath, 1)

	fmt.Println(rootPath)
	PrintDirTree(tree)
}

func PrintDirTree(tree []DirTree) {
	for _, f := range tree {
		padding := fmt.Sprintf("%-*v", f.Depth*2, "")
		if f.IsDir {
			fmt.Println(padding, f.Name)
			PrintDirTree(f.Files)
		} else {
			fmt.Println(padding, f.Name)
		}
	}
}

func BuildDirTree(path string, depth int) (tree []DirTree, err error) {

	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	files := []DirTree{}

	for _, f := range dir {

		subTree := []DirTree{}
		tree := DirTree{
			Name:  f.Name(),
			Path:  path,
			IsDir: f.IsDir(),
			Depth: depth,
			Files: subTree,
		}

		if f.IsDir() {
			tree.Name = filepath.Join(path, tree.Name)
			tree.Name += "/"
			recursiveTree, _ := BuildDirTree(tree.Name, depth+1)
			tree.Files = recursiveTree

		}
		files = append(files, tree)

	}

	return files, nil
}
