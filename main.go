package main

import (
	"os"

	"github.com/volkswagen-onehub/dcc-release/checklist"
)

func main() {
	os.Exit(checklist.Run())
}
