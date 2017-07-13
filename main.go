package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/mlositsky/brbeat/beater"
)

func main() {
	err := beat.Run("brbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
