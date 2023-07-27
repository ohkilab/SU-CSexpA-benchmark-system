package main

import (
	"log"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/cmd/batch/seed"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "batch",
	}
	rootCmd.AddCommand(
		seed.Command,
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
