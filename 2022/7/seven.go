package seven

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
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
		tokens := strings.Split(line, " ")

		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				newPath := string(tokens[2])

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

		if tokens[0] == "dir" {
			if currentDirectory.Directories[tokens[1]] == nil {
				currentDirectory.Directories[tokens[1]] = NewDirectory(tokens[1], currentDirectory)
			}

			// Don't overwrite a directory if it has already been stored
		} else {
			currentDirectory.Files[tokens[1]] = util.ToInt(tokens[0])
		}
	}

	directorySizeCache := map[*Directory]int{}
	directorySizeCache[root] = calculateDirectorySize(root, directorySizeCache)

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

func calculateDirectorySize(current *Directory, cache map[*Directory]int) int {
	totalSize := 0

	for _, fileSize := range current.Files {
		totalSize += fileSize
	}

	for _, directoryPtr := range current.Directories {
		totalSize += calculateDirectorySize(directoryPtr, cache)
	}

	cache[current] = totalSize

	return totalSize
}
