package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

const (
	// path of the JAVA gradle project
	path     = "--PATH--"
	repsPath = "./reps/"
)

var (
	appNameList   []string
	dependencyMap = map[string]int{}
)

func main() {

	// read settings.gradle to catch app names
	readDependencies()

	// read reports of dependensies
	readReps()

	// create output report file
	outFile, err := os.Create("report.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Sort the keys
	keys := make([]string, 0, len(dependencyMap))
	for k := range dependencyMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		_, err := outFile.WriteString(k + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func readReps() {
	files, err := ioutil.ReadDir(repsPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileName := repsPath + f.Name()
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if (strings.Contains(line, "+--- ")) &&
				(!(strings.Contains(line, "|"))) &&
				(!(strings.Contains(line, "(n)"))) &&
				(!(strings.Contains(line, "(*)")) &&
					(!(strings.Contains(line, "project")))) {
				str := strings.Trim(line, "+--- ")
				dependencyMap[str] = dependencyMap[str] + 1
			}
		}
	}
}

func readDependencies() {
	// read settings.gradle to catch app name
	settingsFileName := path + "/settings.gradle"
	settingsFile, err := os.Open(settingsFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer settingsFile.Close()
	scanner := bufio.NewScanner(settingsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "include") {
			str := strings.Trim(line, "include")
			str = strings.Trim(str, " '")
			str = strings.Trim(str, "'")
			appNameList = append(appNameList, str)
		}
	}

	// read libraries and craete report per each app in appNameList
	fmt.Print("Start reading gradle dependencies .")
	for i := 0; i < len(appNameList)-1; i++ {
		fmt.Print(".")
		readGradle(appNameList[i])
	}
	fmt.Println(".")
}

func readGradle(appName string) {
	command := "./gradlew"
	args := appName + ":dependencies"

	// create output file
	outputFileName := "./reps/" + appName + ".txt"
	outfile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	// run command
	cmd := exec.Command(command, args)
	cmd.Dir = path

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(outfile)

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()
}
