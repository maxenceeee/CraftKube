package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"xamence.eu/craftkube/manager/cli"
	"xamence.eu/craftkube/manager/repository"
)

var RepoService *repository.ServiceRepository

// Manager main entry point
func main() {

	RepoService = repository.NewServiceRepository()

	cli.InitCommands(RepoService)

	loopShell(cli.GetRootCommand())
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
