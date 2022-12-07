package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type File struct {
	name string
	parent *Directory
	size int
}

type Directory struct {
	name string
	size int
	files []File
	subdirectories []*Directory
	parent *Directory
}

func main() {
	sampleRoot := getRootForFile("sample_input.txt")
	root := getRootForFile("input.txt")
	answerOne(sampleRoot)
	answerOne(root)

	answerTwo(sampleRoot)
	answerTwo(root)
}

func getRootForFile(filename string) Directory {
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	root := buildDirectoryTree(lines)
	addSizeToDirectories(&root)

	return root
}

func answerTwo(root Directory) {
	totalSpaceAvailable := 70000000
	totalSpaceNeeded := 30000000
	spaceAvailable := totalSpaceAvailable - root.size
	spaceNeeded := totalSpaceNeeded - spaceAvailable

	fmt.Println("Total space available:", spaceAvailable, "Total space needed:", spaceNeeded)

	largeDirs := findDirectoriesWithSizeGreaterThan(&root, spaceNeeded)
	var smallestDir *Directory
	for _, largeDir := range largeDirs {
		fmt.Println((*largeDir).name, (*largeDir).size)
		if smallestDir == nil || (*largeDir).size < (*smallestDir).size {
			smallestDir = largeDir
		}
	}
	fmt.Println("Smallest directory with size greater than", spaceNeeded, "is", (*smallestDir).name, "with size", (*smallestDir).size)
}


func answerOne(root Directory) {
	displayDir(root, 0)
	smallDirs := findDirectoriesWithSizeLessThan(&root, 100000)
	var size int
	for _, smallDir := range smallDirs {
		fmt.Println((*smallDir).name)
		size += (*smallDir).size
	}
	fmt.Println("Total size of directories with size less than 100000:", size)
}

func displayDir(dir Directory, depth int) {
	fmt.Println(strings.Repeat(" ", depth), "-", dir.name, "(dir)", "numFiles", len(dir.files), "numSubdirectories", len(dir.subdirectories), "size", dir.size)
	for _, file := range dir.files {
		fmt.Println(strings.Repeat(" ", depth + 2), "-", file.name, "(file, size=", file.size, ")")
	}
	for _, subdirectory := range dir.subdirectories {
		displayDir(*subdirectory, depth + 2)
	}
}

func findDirectoriesWithSizeGreaterThan(dir *Directory, size int) []*Directory {
	var result []*Directory
	if dir.size > size {
		result = append(result, dir)
	}
	for _, subdirectory := range (*dir).subdirectories {
		result = append(result, findDirectoriesWithSizeGreaterThan(subdirectory, size)...)
	}
	return result
}

func findDirectoriesWithSizeLessThan(dir *Directory, size int) []*Directory {
	var result []*Directory
	if dir.size < size {
		result = append(result, dir)
	}
	for _, subdirectory := range (*dir).subdirectories {
		result = append(result, findDirectoriesWithSizeLessThan(subdirectory, size)...)
	}
	return result
}

func findSubdirectoryWithName(subdirectories []*Directory, name string) *Directory {
	var result *Directory
	for _, subdirectory := range subdirectories {
		if (*subdirectory).name == name {
			result = subdirectory
		}
	}
	return result
}

func addSizeToDirectories(dir *Directory) {
	for _, file := range dir.files {
		dir.size += file.size
	}
	for _, subdirectory := range dir.subdirectories {
		addSizeToDirectories(subdirectory)
		dir.size += (*subdirectory).size
	}
}

func buildDirectoryTree(lines []string) Directory {
	var currentDirectory *Directory
	root := Directory{name: "/", size: 0, files: []File{}, parent: nil}
	for _, line := range lines {
		switch line {
		case "$ cd /":
			currentDirectory = &root
		default:
			commands := strings.Split(line, " ")
			switch commands[0] {
			case "$":
				switch commands[1] {
				case "cd":
					if commands[2] == ".." {
						currentDirectory = (*currentDirectory).parent
					} else {
						foundDirectory := findSubdirectoryWithName(currentDirectory.subdirectories, commands[2])
						currentDirectory = foundDirectory
					}
				}
			case "dir":
				newDirectory := Directory{name: commands[1], size: 0, files: []File{}, parent: currentDirectory}
				currentDirectory.subdirectories = append((*currentDirectory).subdirectories, &newDirectory)
			default:
				size, _ := strconv.Atoi(commands[0])
				currentDirectory.files = append(currentDirectory.files, File{name: commands[1], parent: currentDirectory, size: size})
			}
		}
	}

	return root
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
