package goMetadata

import (
	"context"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello world!")

	uID, uType, err := CheckMetaData(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(uID, uType)
}
