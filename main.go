package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadDirCustom(dirname string, printFiles bool) ([]os.FileInfo, error) {
	var dirsOnly []os.FileInfo
	files, err := ioutil.ReadDir(dirname)
	if !printFiles {
		for _, f := range files {
			if f.IsDir() {
				dirsOnly = append(dirsOnly, f)
			}
		}
		return dirsOnly, err
	}
	return files, err
}

func printFile(prefix string, name string, printFiles, size int64, isDir bool, isLast bool) {
	if isLast {
		graphSymbols := "└───"
	} else {
		graphSymbols := "├───"
	}
	if size == 0 {
		sizeSymbols := "empty"
	} else {
		sizeSymbols := "(%vb)"
	}
	if isDir {
		fmt.Printf("%v%v%v", prefix, graphSymbols, name)
	} else {

		fmt.Printf("%v%v%v", prefix, graphSymbols, name)
	}
}

func dirTreeRecur(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := ReadDirCustom(path, printFiles)
	for index, f := range files {
		if index == len(files)-1 {
			if f.IsDir() {
				fmt.Println(prefix + "└───" + f.Name())
				dirTreeRecur(out, path+string(os.PathSeparator)+f.Name(), printFiles, prefix+"    ")
			} else {
				fmt.Printf("%v└───%v (%vb)\n", prefix, f.Name(), f.Size())
			}
		} else {
			if f.IsDir() {
				fmt.Println(prefix + "├───" + f.Name())
				dirTreeRecur(out, path+string(os.PathSeparator)+f.Name(), printFiles, prefix+"│   ")
			} else {
				fmt.Printf("%v├───%v (%vb)\n", prefix, f.Name(), f.Size())
			}
		}

	}
	return err
}
func dirTree(out io.Writer, path string, printFiles bool) error {
	//fmt.Println(path)
	return dirTreeRecur(out, path, printFiles, "")
}

func main() {
	out := os.Stdout
	path := "."
	if len(os.Args) == 2 || len(os.Args) == 3 {
		//panic("usage go run main.go . [-f]")
		path = os.Args[1]
	}
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
