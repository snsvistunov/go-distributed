package file

import (
	"log"
	"os"
)

func (c *Content) WriteContent(outputDataPath string) {
	f, err := os.Create(outputDataPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, word := range *c {
		_, err := f.WriteString(word + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
