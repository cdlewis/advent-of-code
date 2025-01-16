package nine

import (
	"slices"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
)

type File struct {
	ID         int
	StartingAt int
	Size       int
}

func Nine() int {
	diskMap := aoc.GetInput(9, false, "2333133121414131402")

	var freeBlocks []File
	var disk []File
	nextFileID := 0
	nextBlockPosition := 0
	for idx, i := range diskMap {
		fileSize := int(i - '0')
		fileID := -1

		blockPosition := nextBlockPosition
		nextBlockPosition += fileSize

		if idx%2 == 0 {
			fileID = nextFileID
			nextFileID++
		} else if fileSize == 0 {
			continue // skip 0 empty space blocks
		}

		newFile := File{
			ID:         fileID,
			StartingAt: blockPosition,
			Size:       fileSize,
		}

		if fileID == -1 {
			freeBlocks = append(freeBlocks, newFile)
		} else {
			disk = append(disk, newFile)
		}
	}

	for i := len(disk) - 1; i >= 0; i-- {
		for j := 0; j < len(freeBlocks); j++ {
			if freeBlocks[j].StartingAt > disk[i].StartingAt {
				break
			}

			if freeBlocks[j].Size >= disk[i].Size {
				disk[i].StartingAt = freeBlocks[j].StartingAt
				freeBlocks[j].StartingAt += disk[i].Size
				freeBlocks[j].Size -= disk[i].Size

				break
			}
		}
	}

	checkSum := 0
	slices.SortFunc(disk, func(i, j File) int {
		return i.StartingAt - j.StartingAt
	})

	for _, file := range disk {
		for j := range file.Size {
			checkSum += (file.StartingAt + j) * file.ID
		}

	}

	return checkSum
}
