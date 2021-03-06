package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Originate/exosphere/src/application"
	"github.com/Originate/exosphere/src/config"
	"github.com/Originate/exosphere/src/docker/composebuilder"
	"github.com/Originate/exosphere/src/types"
	"github.com/Originate/exosphere/src/util"
	"github.com/spf13/cobra"
)

var deployProfileFlag string

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys Exosphere application to the cloud",
	Long:  "Deploys Exosphere application to the cloud",
	Run: func(cmd *cobra.Command, args []string) {
		if printHelpIfNecessary(cmd, args) {
			return
		}
		fmt.Println("We are about to deploy an application!")
		appDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		homeDir, err := util.GetHomeDirectory()
		if err != nil {
			panic(err)
		}
		appConfig, err := types.NewAppConfig(appDir)
		if err != nil {
			log.Fatalf("Cannot read application configuration: %s", err)
		}
		serviceConfigs, err := config.GetServiceConfigs(appDir, appConfig)
		if err != nil {
			log.Fatalf("Failed to read service configurations: %s", err)
		}
		awsConfig, err := getAwsConfig(deployProfileFlag)
		if err != nil {
			log.Fatalf("Failed to read secrest configurations: %s", err)
		}

		logger := application.NewLogger([]string{"exo-deploy"}, []string{}, os.Stdout)
		terraformDir := filepath.Join(appDir, "terraform")
		deployConfig := types.DeployConfig{
			AppConfig:                appConfig,
			ServiceConfigs:           serviceConfigs,
			AppDir:                   appDir,
			HomeDir:                  homeDir,
			DockerComposeProjectName: composebuilder.GetDockerComposeProjectName(appDir),
			LogChannel:               logger.GetLogChannel("exo-deploy"),
			TerraformDir:             terraformDir,
			SecretsPath:              filepath.Join(terraformDir, "secrets.tfvars"),
			AwsConfig:                awsConfig,
		}

		err = application.StartDeploy(deployConfig)
		if err != nil {
			log.Fatalf("Deploy failed: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().StringVarP(&deployProfileFlag, "profile", "p", "default", "AWS profile to use")
}
