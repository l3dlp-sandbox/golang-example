package xmlToCsv

import (
	"fmt"
	"os"
)

func XmlToCsv() {
	xmlFile, err := os.Open("/home/damitha/xmlCode/Posts1.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(xmlFile)
}
