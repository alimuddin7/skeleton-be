package main

import (
	"context"
	"os"

	"andromeda.ottopay.id/pt-rtsm-ottopay/skeleton-svc/cmd"
	"github.com/charmbracelet/fang"
)

func main() {
	if err := fang.Execute(context.Background(), cmd.Root()); err != nil {
		os.Exit(1)
	}
}
