package csvFile
 
import (
	"encoding/csv"
	"bufio"
	"fmt"
	"log"
	"os"
)
 
/*
Reading CSV File
The csv package have a NewReader() function which returns a Reader object to process CSV data. A csv.Reader converts \r\n sequences in its input to just \n, which includes multi line field values also.

The file test.csv have few records is opened in read-only mode using the os.Open() function, which returns an pointer type instance of os.File. The csv.Reader.Read() method is used to decode each file record into pre-defined struct CSVData and then store them in a slice until io.EOF is returned.
*/
func readCsv() {
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
Writing CSV File
The csv package have a NewWriter() function which returns a Writer object which is used for writing CSV data. A csv.Writer writes csv records which are terminated by a newline and uses a comma as the field delimiter. The following source code snippet shows how to write data to a CSV file.

A two-dimensional slice rows contains sample csv records. The os.Create() function creates a csv file test.csv; truncate all it's records if already exists and returning an instance of os.File object. The csvwriter.Write(row) method is called to write each slice of strings to the file as CSV records.
*/

func writeCsv() {
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Pinky", "London", "Python"},
		{"Nicky", "Paris", "Golang"},
		{"Micky", "Tokyo", "Php"},
	}
 
	csvfile, err := os.Create("test.csv")
 
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
 
	csvwriter := csv.NewWriter(csvfile)
 
	for _, row := range rows {
		_ = csvwriter.Write(row)
	}
 
	csvwriter.Flush()
 
	csvfile.Close()
}
