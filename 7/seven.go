package seven

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

type Directory struct {
	Directories map[string]*Directory
	Files       map[string]int
	Name        string
	Parent      *Directory
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:        name,
		Files:       map[string]int{},
		Directories: map[string]*Directory{},
		Parent:      parent,
	}
}

func Seven(useExample bool, example string) int {
	raw := util.GetInput(7, useExample, example)

	root := NewDirectory("/", nil)
	currentDirectory := root

	for _, line := range strings.Split(raw, "\n") {
		if line[0] == '$' {
			if string(line[:4]) == "$ cd" {
				newPath := string(line[5:])

				if newPath == "/" {
					currentDirectory = root
				} else if newPath == ".." {
					currentDirectory = currentDirectory.Parent
				} else {
					currentDirectory = currentDirectory.Directories[newPath]
				}
			}

			// We can ignore the ls command

			continue
		}

		parts := strings.Split(line, " ")

		if parts[0] == "dir" {
			if currentDirectory.Directories[parts[1]] == nil {
				currentDirectory.Directories[parts[1]] = NewDirectory(parts[1], currentDirectory)
			}
		} else {
			currentDirectory.Files[parts[1]] = util.ToInt(parts[0])
		}
	}

	directorySizeCache := map[*Directory]int{}

	// Recursively calculate directory sizes
	var calculateDirectorySize func(current *Directory) int
	calculateDirectorySize = func(current *Directory) int {
		totalSize := 0

		for _, fileSize := range current.Files {
			totalSize += fileSize
		}

		for _, directoryPtr := range current.Directories {
			totalSize += calculateDirectorySize(directoryPtr)
		}

		directorySizeCache[current] = totalSize

		return totalSize
	}
	directorySizeCache[root] = calculateDirectorySize(root)

	currentFree := 70000000 - directorySizeCache[root]
	spaceNeeded := 30000000 - currentFree

	fmt.Println("Searching for best candidate with", spaceNeeded, "required")

	winningSize := directorySizeCache[root]
	for _, directorySize := range directorySizeCache {
		if directorySize > spaceNeeded && directorySize < winningSize {
			winningSize = directorySize
		}
	}

	return winningSize
}
