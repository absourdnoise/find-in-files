package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"runtime"
	"os"
	"strings"
	"path/filepath"
)

var fileTypes []string
var findTo = "getMenu13"
var finded map[int]string
var MAXPROC int
var rootDir = "/home/andrew/work/_NEW_LS_/templates/"

func main()	{
	fileTypes = []string{".html", ".asp"}
	var tasks []string

	MAXPROC = runtime.NumCPU() - 1

	dirf, _ := ioutil.ReadDir(rootDir)


	for i, f := range dirf {
		if f.IsDir() {
			tasks = append(tasks, f.Name())
			if i > 0 && (i % MAXPROC) == 0 {
				go doSearchInSlice(tasks);
				doSearchInSlice(tasks);
				tasks = nil;
			}
		}
	}
	go doSearchInSlice(tasks);
	doSearchInSlice(tasks);
}

func searchInFile(findTo string, file string) (results map[int]string, err error) {
	err = nil
	results = make(map[int]string)

	inFile, err := os.Open(file)

	defer inFile.Close()

	fileContent := bufio.NewReader(inFile)
	eof := false;
	for i := 1; !eof; i++ {
		var fileLine string;
		fileLine, err = fileContent.ReadString('\n');
		if err == io.EOF {
			err = nil   // io.EOF isn't really an error
			eof = true  // this will end the loop at the next iteration
		} else if err != nil {
			return nil, err  // finish immediately for real errors
		}

		if strings.Contains(strings.ToLower(fileLine), strings.ToLower(findTo)) {
			results[i] = fileLine;
			fmt.Printf("%s:%d: %s", file, i, fileLine);
		}
	}

	return results, err
}

func scanFs(path string, f os.FileInfo, err error) error {
	for _, ft := range fileTypes{
		idx := len(path) - len(ft)

		if idx >= 0 && path[idx:] == ft {
			searchInFile(findTo, path);
		}
	}
	return nil
}

func doSearchInSlice(dirs []string) error {
	for _, f := range dirs {
		filepath.Walk(rootDir + f + "/templates", scanFs);
	}
	return nil
}
