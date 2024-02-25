package docxtext

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
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

			content, err := io.ReadAll(rc) // Changed from ioutil.ReadAll to io.ReadAll
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
	var resultStrings []string
	var builder strings.Builder // 使用strings.Builder优化字符串拼接
	inParagraph := false

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
			if se.Name.Local == "p" {
				inParagraph = true
				builder.Reset() // 重置builder准备收集新段落
			}
		case xml.CharData:
			if inParagraph {
				builder.WriteString(string(se)) // 收集段落文本
			}
		case xml.EndElement:
			if se.Name.Local == "p" && builder.Len() > 0 {
				resultStrings = append(resultStrings, builder.String())
				inParagraph = false // 更新状态，表示当前不在段落内
			}
		}
	}

	return resultStrings, nil
}
