package document

import (
	"encoding/xml"
	"fmt"
	doc_xml "github.com/OntoLedgy/storage_interop_services/code/services/documents/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestXSDIO(t *testing.T) {
	sourceFileName := "D:\\Temp\\2021-03-04_lei-cdf-v3-1.xsd"
	TargetFileName := "D:\\Temp\\2021-03-04_lei-cdf-v3-1.go"
	TargetLanguage := "Go"

	doc_xml.ConvertXSDtoCode(
		sourceFileName,
		TargetFileName,
		TargetLanguage)

}

//TODO - wrap and move this to main code
func TestXMLIO(t *testing.T) {

	// Open our xmlFile
	xmlFile, err := os.Open("D:\\Temp\\20221030-0800-gleif-goldencopy-lei2-golden-copy.xml")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var leiData LEIData

	xml.Unmarshal(byteValue, &leiData)

	for i := 0; i < 100; i++ {
		leiRecord := leiData.LEIRecords[0]
		fmt.Println("Record: " + leiRecord.LEIRecord[i].Entity.LegalName.Value)

	}
}
