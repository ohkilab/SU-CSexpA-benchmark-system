package main

import (
	"log"

	create_users "github.com/ohkilab/SU-CSexpA-benchmark-system/backend/cmd/batch/create-users"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/cmd/batch/seed"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "batch",
	}
	rootCmd.AddCommand(
		seed.Command,
		create_users.Command,
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
