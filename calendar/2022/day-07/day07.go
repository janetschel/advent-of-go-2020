package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/sets"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type directory struct {
	name string
	files map[string]int
	nestedDirectoryNames *sets.Set
	parentDirectory string
}

var smallestDirectoryLimit, totalDiskSpace, necessaryFreeSpace = 100000, 70000000, 30000000

func main() {
	input := files.ReadFile(7, 2022, "\n")
	directories := parseInput(input)
	sumOfSmallestDirectories, smallestDirectoryToDelete := analyzeDirectories(directories)
	println(sumOfSmallestDirectories)
	println(smallestDirectoryToDelete)
}

func analyzeDirectories(directories map[string]directory) (int, int) {
	smallestSum, smallestToDelete := 0, math.MaxInt
	totalUsed := getDirectorySize(directories["/"], directories)
	for _, dir := range directories {
		totalSize := getDirectorySize(dir, directories)		
		if totalSize <= smallestDirectoryLimit {
			smallestSum += totalSize
		}
		if totalSize < smallestToDelete && totalDiskSpace - (totalUsed - totalSize) >= necessaryFreeSpace {
			smallestToDelete = totalSize
		}
	}

	return smallestSum, smallestToDelete
}

func parseInput(input []string) map[string]directory {
	directories := map[string]directory{}

	currentDirName := ""
	for _, output := range input {
		if output == "$ cd .." {
			currentDirName = directories[currentDirName].parentDirectory
		} else if len(output) >= 5 && output[0:5] == "$ cd " {
			suffix := "/"
			dirBaseName := strings.Fields(output)[2]
			if dirBaseName == suffix {
				suffix = ""
			}
			currentDirName += dirBaseName + suffix
			_, hasDir := directories[currentDirName]
			if !hasDir {
				newSet := sets.New()
				directories[currentDirName] = directory{name: currentDirName, nestedDirectoryNames: &newSet, files: map[string]int{} }
			}
		} else if len(output) >= 4 && output[0:4] == "dir " {
			dirName := currentDirName + strings.Fields(output)[1] + "/"
			dir, hasDir := directories[dirName]
			if !hasDir {
				newSet := sets.New()
				directories[dirName] = directory{name: currentDirName, nestedDirectoryNames: &newSet, parentDirectory: currentDirName, files: map[string]int{} }
			} else {
				dir.parentDirectory = currentDirName
			}
			directories[currentDirName].nestedDirectoryNames.Add(dirName)
		} else if output != "$ ls" {
			// output line is a file name and size
			size, name := parseFile(output)
			fileSize, hasFile := directories[currentDirName].files[name]
			if !hasFile {
				directories[currentDirName].files[name] = size
			} else {
				panic(fmt.Sprintf("already processed file %s in dir %s, size %d vs existing size %d", name, currentDirName, size, fileSize))
			}
		}
	}

	return directories
}

func parseFile(output string) (int, string) {
	parts := strings.Fields(output)
	size, _ := strconv.Atoi(parts[0])
	return size, parts[1]
}

func getDirectorySize(dir directory, allDirectories map[string]directory) int {
	topLevelSize := 0
	for _, fileSize := range dir.files {
		topLevelSize += fileSize
	}
	for _, childDirName := range dir.nestedDirectoryNames.Iterator() {
		childDir := allDirectories[childDirName] 
		topLevelSize += getDirectorySize(childDir, allDirectories)
	}
	return topLevelSize
}
