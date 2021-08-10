package uuid

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
)

func Run(args []string) {
	numberOfArgs := len(args)
	if numberOfArgs == 0 {
		fmt.Printf("Randomly generated UUID: %s\n", uuid.New())
	} else {
		for i := 0; i < numberOfArgs; i++ {
			uuid, err := uuid.Parse(args[i])
			if err != nil {
				fmt.Printf("%s string parsed with error %s\n", humanize.Ordinal(i+1), err)
			} else {
				fmt.Printf("%s string parsed successfully: %s\n", humanize.Ordinal(i+1), uuid)
			}
		}
	}

	fmt.Println("Parsing successfully")
}
