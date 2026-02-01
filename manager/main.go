package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"xamence.eu/craftkube/manager/parser"
	"xamence.eu/craftkube/manager/repository"
)

var repo *repository.ServiceRepository

// Manager main entry point
func main() {

	repo = repository.NewServiceRepository()

	var serviceCmd = &cobra.Command{
		Use: "service",
		Run: createServiceCmd,
	}

	serviceCmd.Flags().StringP("file", "f", "", "Path to the service YAML file")
	serviceCmd.MarkFlagRequired("file")

	var rootCmd = &cobra.Command{
		Use: "craftkube",
	}

	rootCmd.AddCommand(serviceCmd)

	loopShell(rootCmd)

}

func createServiceCmd(cmd *cobra.Command, args []string) {
	// Implementation of the service command
	filePath := cmd.Flag("file").Value.String()
	service, err := parser.ParseServiceYAMLFile(filePath)
	if err != nil {
		panic(err)
	}

	repo.AddService(*service)
}

func loopShell(rootCmd *cobra.Command) {
	// Implementation of the shell loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("CraftKube> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "exit" || line == "quit" {
			break
		}
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		rootCmd.SetArgs(args)
		rootCmd.Execute()
	}
}
