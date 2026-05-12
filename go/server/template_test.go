package server

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

func TestAllTemplatesParse(t *testing.T) {
	templateDir := "../templates"
	tmplCount := 0
	failCount := 0
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, err := template.ParseFiles(path)
			if err != nil {
				t.Errorf("FAIL: Template parse error in %s: %v", path, err)
				failCount++
			} else {
				t.Logf("PASS: %s", path)
				tmplCount++
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Error walking template directory: %v", err)
	}
	if tmplCount == 0 {
		t.Fatalf("No templates found in %s", templateDir)
	}
	if failCount > 0 {
		t.Fatalf("%d template(s) failed to parse", failCount)
	} else {
		t.Logf("All %d templates parsed successfully.", tmplCount)
		fmt.Println("TEMPLATE TEST: PASS")
	}
}
