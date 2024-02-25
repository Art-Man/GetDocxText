package docxtext

import (
	"log"
	"path/filepath"
	"testing"
)

// TestGetXmlContent tests the extraction of XML content from a .docx file.
func TestGetXmlContent(t *testing.T) {
	// Define the path to the test .docx file
	docxFilePath := filepath.Join("testdata", "example.docx")

	// Call GetXmlContent
	content, err := GetXmlContent(docxFilePath)
	if err != nil {
		t.Errorf("GetXmlContent returned an error: %v", err)
	}

	// Check if content is not nil (a simple validation)
	if content == nil || len(content) == 0 {
		t.Errorf("Expected XML content to be non-empty, got nil or empty")
	} else {
		log.Println(string(content))
	}
}

// TestGetTextByParagraph tests the parsing of XML content to extract text by paragraphs,
// specifically looking for "My resume" and "Personal Skills" in the extracted text.
func TestGetTextByParagraph(t *testing.T) {
	// Since this depends on the output of GetXmlContent, start by extracting the content.
	docxFilePath := filepath.Join("testdata", "example.docx")
	content, err := GetXmlContent(docxFilePath)
	if err != nil {
		t.Fatalf("Failed to extract XML content: %v", err)
	}

	// Call GetTextByParagraph with the extracted content.
	texts, err := GetTextByParagraph(content)
	if err != nil {
		t.Errorf("GetTextByParagraph returned an error: %v", err)
	}

	// Check for the presence of specific paragraphs.
	expectedTexts := []string{"My resume", "Personal Skills"}
	foundTexts := make(map[string]bool)
	for _, expected := range expectedTexts {
		foundTexts[expected] = false // Initialize map with expected texts set to false
	}

	// Iterate over the extracted texts to check if the expected texts are present.
	for _, text := range texts {
		if _, ok := foundTexts[text]; ok {
			foundTexts[text] = true
		}
	}

	// Verify that each expected text was found.
	for _, expected := range expectedTexts {
		if !foundTexts[expected] {
			t.Errorf("Expected text %q not found in paragraphs", expected)
		}
	}

	// Optionally log found texts for verification.
	for _, text := range texts {
		log.Println(text)
	}
}
