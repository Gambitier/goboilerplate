package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gambitier/gocomm/config"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "scripts",
		Short: "A CLI application for executing scripts",
	}

	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Manage database migrations",
	}

	migrateCreateCmd := &cobra.Command{
		Use:   "create [MigrationName]",
		Short: "Create a new database migration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			migrationName := args[0]
			createMigration(migrationName)
		},
	}

	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "Run all pending database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			runMigrations()
		},
	}

	migrateCmd.AddCommand(migrateCreateCmd, migrateUpCmd)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func createMigration(migrationName string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "db/migrations", "-seq", migrationName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error executing migrate create: %v", err)
	}
	fmt.Println("Migration created successfully.")
}

func runMigrations() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	cmd := exec.Command("migrate", "-database", cfg.DatabaseURL, "-path", "db/migrations", "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error executing migrations: %v", err)
	}
	fmt.Println("Migrations executed successfully.")
}
