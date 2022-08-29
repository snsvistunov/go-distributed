package file

import (
	"log"
	"os"
	"strings"
)

type Content []string

func (c *Content) NewContent(dataPath string) {
	s, err := os.ReadFile(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	*c = strings.Fields(string(s))
}
