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
//var finded map[int]string
var MAXPROC int
var rootDir = "/home/andrew/work/_NEW_LS_/templates/"
var queue = "/usr/wwwc/_main-ef-panel/includes/squeue/"

func main()	{
	fileTypes = []string{".html", ".asp"}
	var tasks []string
	var prefix = "/templates"

	rootDir, prefix, findTo, _ = filenamesFromCommandLine();

	MAXPROC = runtime.NumCPU() - 1
	dirf, _ := ioutil.ReadDir(rootDir)

	for i, f := range dirf {
		if f.IsDir() {
			tasks = append(tasks, f.Name())
			if i > 0 && (i % MAXPROC) == 0 {
				go doSearchInSlice(tasks, prefix);
				doSearchInSlice(tasks, prefix);
				tasks = nil;
			}
		}
	}
	go doSearchInSlice(tasks, prefix);
	doSearchInSlice(tasks, prefix);
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
			err = nil
			eof = true
		} else if err != nil {
			return nil, err
		}

		if strings.Contains(strings.ToLower(fileLine), strings.ToLower(findTo)) {
			results[i] = fileLine;
			fmt.Printf("%s : %s", file, fileLine);
		}
	}

	return results, err
}

func scanFs(path string, f os.FileInfo, err error) error {
	for _, ft := range fileTypes {
		idx := len(path) - len(ft)

		if idx >= 0 && path[idx:] == ft {
			searchInFile(findTo, path);
		}
	}
	return nil
}

func doSearchInSlice(dirs []string, prefix string) error {
	for _, f := range dirs {
		filepath.Walk(rootDir + f + prefix, scanFs);
	}
	return nil
}

func filenamesFromCommandLine() (rootdir, subdir, toFind string,
	err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s rootdir hash subdir",
			filepath.Base(os.Args[0]))
		return "", "", "", err
	}
	if len(os.Args) > 1 {
		rootdir = os.Args[1]
		if len(os.Args) > 2 {
			toFind = getFindString(os.Args[2])
			if len(os.Args) > 3 {
				subdir = os.Args[3]
			}
		}
	}

	return rootdir, subdir, toFind, nil
}

func getFindString(hash string) (toFind string) {
	inFile, _ := os.Open(queue + hash)
	defer inFile.Close()

	fileContent := bufio.NewReader(inFile)
	toFind, _ = strings.TrimSpace(fileContent.ReadString('\n'));

	return toFind
}
