package cmd

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/Originate/exosphere/exo-go/src/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/moby/moby/client"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes dangling Docker images and volumes",
	Run: func(cmd *cobra.Command, args []string) {
		if printHelpIfNecessary(cmd, args) {
			return
		}
		fmt.Print("We are about to clean up your Docker workspace!\n\n")
		c, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		options := types.ImageRemoveOptions{
			Force:         true,
			PruneChildren: true,
		}
		cwd, _ := os.Getwd()
		dockerComposePath := path.Join(cwd, "tmp", "docker-compose.yml")
		dockerCompose, err := docker.GetDockerCompose(dockerComposePath)
		if err != nil {
			panic(err)
		}
		if &dockerCompose == nil {
			fmt.Println("Completed, no docker images to remove...")
		}
		for serviceName := range dockerCompose.Services {
			_, err = c.ImageRemove(context.Background(), serviceName, options)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("removed all exosphere related images")
		_, err = c.ImagesPrune(context.Background(), filters.NewArgs())
		if err != nil {
			panic(err)
		}
		fmt.Println("removed all dangling images")
		_, err = c.VolumesPrune(context.Background(), filters.NewArgs())
		if err != nil {
			panic(err)
		}
		fmt.Println("removed all dangling volumes")
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
}
