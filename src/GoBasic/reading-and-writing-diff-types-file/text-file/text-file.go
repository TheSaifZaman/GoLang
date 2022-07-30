package textFile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
 
/*
Reading Text File
The bufio package Scanner generally used for reading the text by lines or words from a file. The following source code snippet shows reading text line-by-line from the plain text file as below.

The os.Open() function is used to open a specific text file in read-only mode and this returns a pointer of type os.File. The method os.File.Close() is called on the os.File object to close the file and there is a loop to iterates through and prints each of the slice values. The program after execution shows the below output line-by-line as they read it from the file.
*/
func readText() {
	file, err := os.Open("test.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}

/*
Writing Text File
The bufio package provides an efficient buffered Writer which queues up bytes until a threshold is reached and then finishes the write operation to a file with minimum resources. The following source code snippet shows writing a string slice to a plain text file line-by-line.

The sampledata is represented as a string slice which holds few lines of data which will be written to a new line within the file. The function os.OpenFile() is used with a flag combination to create a write-only file if none exists and appends to the file when writing.
*/

func writeText() {
	sampledata := []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Nunc a mi dapibus, faucibus mauris eu, fermentum ligula.",
		"Donec in mauris ut justo eleifend dapibus.",
		"Donec eu erat sit amet velit auctor tempus id eget mauris.",
	}
 
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	datawriter := bufio.NewWriter(file)
 
	for _, data := range sampledata {
		_, _ = datawriter.WriteString(data + "\n")
	}
 
	datawriter.Flush()
	file.Close()
}

