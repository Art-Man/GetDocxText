package docxtext

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

// GetXmlContent extracts the XML content from a .docx file.
func GetXmlContent(docxFile string) ([]byte, error) {
	r, err := zip.OpenReader(docxFile)
	if err != nil {
		return nil, fmt.Errorf("error opening docx file: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				return nil, fmt.Errorf("error opening document.xml: %w", err)
			}
			defer rc.Close()

			content, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, fmt.Errorf("error reading document.xml: %w", err)
			}

			return content, nil
		}
	}

	return nil, fmt.Errorf("document.xml not found")
}

// GetTextByParagraph parses XML content and extracts text by paragraphs.
func GetTextByParagraph(content []byte) ([]string, error) {
	decoder := xml.NewDecoder(bytes.NewReader(content))
	var inElement string
	var paragraphText string
	var resultStrings []string

	for {
		t, err := decoder.Token()
		if t == nil {
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("error decoding XML: %w", err)
		}

		switch se := t.(type) {
		case xml.StartElement:
			inElement = se.Name.Local
			if inElement == "p" {
				paragraphText = ""
			}
		case xml.CharData:
			if inElement == "t" {
				paragraphText += string(se)
			}
		case xml.EndElement:
			if se.Name.Local == "p" && len(paragraphText) > 0 {
				resultStrings = append(resultStrings, paragraphText)
				paragraphText = ""
			}
			inElement = ""
		}
	}

	return resultStrings, nil
}
