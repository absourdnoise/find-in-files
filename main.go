package main

import (
	"bufio"
	"fmt"
	"io"
//	"io/ioutil"
//	"log"
	"os"
//	"path/filepath"
//	"regexp"
	"strings"
	"path/filepath"
)

var fileTypes []string
var findTo = "getMenu13"
var finded map[int]string

func main()	{
	fileTypes = []string{".html", ".asp"}

	filepath.Walk("/home/andrew/work/_NEW_LS_/templates/b6-black/templates", scanFs);
	/*var finded map[int]string
	finded, err := searchInFile(findTo, fileToread);
	searchInFile(findTo, fileToread);
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(finded) > 0 {
		for num, str := range finded {
			fmt.Printf("%d: %s", num, str)
		}
	} else {
		fmt.Printf("Not found")
	}*/
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

//		fmt.Println(path, ft, path[idx:])
		if idx >= 0 && path[idx:] == ft {
//			fmt.Println(path)
//			finded, err := searchInFile(findTo, path);
			go searchInFile(findTo, path);
			searchInFile(findTo, path);
	/*		if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
*/
/*			if len(finded) > 0 {
				for num, str := range finded {
					fmt.Printf("%d: %s", num, str)
				}
			}*/
		}
	}

	return nil
}
