package src

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// Get full path from a given string
// Example: func GetFullPath("$HOME/.locale/share/psycofolio2")
// Result: /home/user/.locale/share/psycofolio2
func GetFullPath(path string) (string, error) {
	re := regexp.MustCompile(`\$\w*`)
	strSlice := re.FindAllString(path, -1)
	for _, w := range strSlice {
		envVal := os.Getenv(strings.Replace(w, "$", "", -1))
		if envVal == "" {
			return "", errors.New("Environment Variable \"" + w + "\" is not set on your system")
		}
		path = strings.Replace(path, w, envVal, -1)
	}
	return path, nil
}
