package main

import (
	"context"
	"os"

	"github.com/alimuddin7/skeleton-be/cmd"
	"github.com/charmbracelet/fang"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.Root()); err != nil {
		os.Exit(1)
	}
}
