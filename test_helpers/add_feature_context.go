package testHelpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/pkg/errors"
)

// AddFeatureContext defines the festure context for features/add/**/*.feature
// nolint gocyclo
func AddFeatureContext(s *godog.Suite) {
	s.Step(`^my application contains the template folder "([^"]*)" with the files:$`, func(dirPath string, table *gherkin.DataTable) error {
		if err := os.MkdirAll(path.Join(appDir, dirPath), 0777); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to create %s", dirPath))
		}
		for _, row := range table.Rows[1:] {
			file, content := row.Cells[0].Value, row.Cells[1].Value
			match, err := regexp.MatchString("/", file)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("Failed to parse %s", file))
			}
			if match {
				if err := os.MkdirAll(path.Join(appDir, dirPath, filepath.Dir(file)), 0777); err != nil {
					return errors.Wrap(err, fmt.Sprintf("Failed to create the necessary directories for %s", file))
				}
			}
			if err := ioutil.WriteFile(path.Join(appDir, dirPath, file), []byte(content), 0777); err != nil {
				return errors.Wrap(err, fmt.Sprintf("Failed to create %s", file))
			}
		}
		return nil
	})

	s.Step(`^my application now contains the file "([^"]*)" containing the text:$`, func(fileName string, expectedContent *gherkin.DocString) error {
		bytes, err := ioutil.ReadFile(path.Join(appDir, fileName))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to read %s", fileName))
		}
		return validateTextContains(strings.TrimSpace(string(bytes)), strings.TrimSpace(expectedContent.Content))
	})
}
