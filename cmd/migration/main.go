package main

import (
	"fmt"

	"github.com/TrinhTrungDung/togo/internal/migration"
)

func main() {
	if err := migration.Run(); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}
}
