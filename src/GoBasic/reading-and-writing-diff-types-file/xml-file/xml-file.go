package xmlFile

/*
<note>
<to>Tove</to>
<from>Jani</from>
<heading>Reminder</heading>
<body>Don't forget me this weekend!</body>
</note>
*/
 
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)
 
type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}
 
/*
Reading XML file
The xml package includes Unmarshal() function that supports decoding data from a byte slice into values. The xml.Unmarshal() function is used to decode the values from the XML formatted file into a Notes struct.
Sample XML file:

The notes.xml file is read with the ioutil.ReadFile() function and a byte slice is returned, which is then decoded into a struct instance with the xml.Unmarshal() function. The struct instance member values are used to print the decoded data.
*/
func ReadXml() {
	data, _ := ioutil.ReadFile("notes.xml")
 
	note := &Notes{}
 
	_ = xml.Unmarshal([]byte(data), &note)
 
	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
}

/*
Writing XML file
The xml package has an Marshal() function which is used to serialized values from a struct and write them to a file in XML format.

The notes struct is defined with an uppercase first letter and ″xml″ field tags are used to identify the keys. The struct values are initialized and then serialize with the xml.Marshal() function. The serialized XML formatted byte slice is received which then written to a file using the ioutil.WriteFile() function.
*/

func WriteXml() {
	note := Notes{To: "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}
 
	file, _ := xml.MarshalIndent(note, "", " ")
 
	_ = ioutil.WriteFile("notes1.xml", file, 0644)
 
}