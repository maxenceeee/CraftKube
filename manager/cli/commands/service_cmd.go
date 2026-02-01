package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"xamence.eu/craftkube/manager/parser"
	"xamence.eu/craftkube/manager/repository"
)

type ServiceCommand struct {
	RepoService *repository.ServiceRepository
}

func (s *ServiceCommand) Execute(c *cobra.Command, args []string) error {

	// Implementation of the service command

	filePath := c.Flag("file").Value.String()

	service, err := parser.ParseServiceYAMLFile(filePath)
	if err != nil {
		fmt.Printf("Error parsing YAML file (location: %s): %v\n", filePath, err)
		return err
	}

	s.RepoService.AddService(*service)

	return nil
}

func (s *ServiceCommand) Name() string {
	return "service"
}

func (s *ServiceCommand) Description() string {
	return "Manage services"
}

func (s *ServiceCommand) Flags() map[string]string {
	return map[string]string{
		"file": "Path to the service YAML file",
	}
}
