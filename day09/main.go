package main

import (
	"fmt"
	"strconv"

	_ "embed"
)

//go:embed input.txt
var input string

type Node struct {
	ID   int
	Pos  int
	Free bool
	Len  int
}

type FileSystem []Node

func buildFileSystem(fsMap string) FileSystem {
	fileSystem := FileSystem{}
	fileID := 0
	for index, file := range fsMap {
		node := Node{ID: -1, Pos: index, Free: true}
		length, _ := strconv.Atoi(string(file))
		if index%2 == 0 {
			node.ID = fileID
			node.Free = false
			node.Len = length
			fileID++
		} else {
			node.Len = length
		}

		for range length {
			fileSystem = append(fileSystem, node)
		}
	}
	return fileSystem
}

func calcChecksum(fileSystem FileSystem) int {
	checksum := 0
	for pos, node := range fileSystem {
		if !node.Free {
			checksum += node.ID * pos
		}
	}
	return checksum
}

func part1(input string) int {
	fileSystem := buildFileSystem(input)

	lastIndex := len(fileSystem) - 1
	for index := 0; index < lastIndex; index++ {
		node := fileSystem[index]
		if node.Free {
			fileSystem[index].Free = false
			fileSystem[index].ID = fileSystem[lastIndex].ID
			fileSystem[lastIndex].Free = true
			lastIndex--
		}
		for fileSystem[lastIndex].Free == true {
			lastIndex--
		}
	}

	return calcChecksum(fileSystem)
}

func part2(input string) int {
	fileSystem := buildFileSystem(input)
	for index := range fileSystem {
		for lastIndex := len(fileSystem) - 1; lastIndex > 0; lastIndex-- {
			node := fileSystem[index]
			lastNode := fileSystem[lastIndex]
			if index > lastIndex {
				break
			}

			if node.Free && !lastNode.Free && node.Len >= lastNode.Len {
				i := 0
				for ; i < lastNode.Len; i++ {
					fileSystem[index+i].Free = false
					fileSystem[index+i].ID = lastNode.ID
					fileSystem[lastIndex].Free = true
					lastIndex--
				}
				for ; i < node.Len; i++ {
					fileSystem[index+i].Len = node.Len - lastNode.Len
				}
			}
		}
	}
	return calcChecksum(fileSystem)
}

func main() {
	part1Result := part1(input)
	fmt.Printf("Part 1 total: %d\n", part1Result)

	part2Result := part2(input)
	fmt.Printf("Part 2 total: %d\n", part2Result)
}
