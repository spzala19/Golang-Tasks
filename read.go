package main

import (
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	//Place names.txt in same directory (format: firstname lastname)
	// opening a file
	fmt.Println("Enter file name which consists names:")
	var fileName string
	fmt.Scan(&fileName)
	f, error := os.Open(fileName)
	//finding file size
	fi, statErr := f.Stat()
	var content []byte
	if statErr == nil && error == nil {
		content = make([]byte, fi.Size())
		//reading file content
		nb, readErr := f.Read(content)
		if readErr == nil && nb > 0 {
			//fmt.Print(string(content), "\n")
		} else {
			fmt.Print(readErr)
		}
	} else {
		fmt.Print(error)
	}
	//Core logic
	strs := strings.Split((string(content)), "\n")
	var n name
	var nstr []string
	sli := make([]name, len(strs), len(strs)+1)
	for i := 0; i < len(strs); i++ {

		nstr = strings.Split(strs[i], " ")
		n.fname = nstr[0]
		n.lname = nstr[1]
		sli[i] = n

	}
	// Printing fname and lname from slice of structs
	for i := 0; i < len(sli); i++ {
		fmt.Println("Fname = :", sli[i].fname)
		fmt.Println("Lname = :", sli[i].lname)
	}
}
