package config

import (
	"fmt"

	"github.com/Originate/exosphere/src/docker/tools"
	"github.com/Originate/exosphere/src/types"
)

type genericDependency struct {
	config    types.DependencyConfig
	appConfig types.AppConfig
	appDir    string
	homeDir   string
}

// GetContainerName returns the container name
func (g *genericDependency) GetContainerName() string {
	return g.config.Name + g.config.Version
}

//GetDeploymentConfig returns configuration needed in deployment
func (g *genericDependency) GetDeploymentConfig() (map[string]string, error) {
	config := map[string]string{
		"version": g.config.Version,
	}
	return config, nil
}

// GetDeploymentServiceEnvVariables returns configuration needed for each service in deployment
func (g *genericDependency) GetDeploymentServiceEnvVariables() map[string]string {
	return map[string]string{}
}

// GetDockerConfig returns docker configuration and an error if any
func (g *genericDependency) GetDockerConfig() (types.DockerConfig, error) {
	renderedVolumes, err := tools.GetRenderedVolumes(g.config.Config.Volumes, g.appConfig.Name, g.config.Name, g.homeDir)
	if err != nil {
		return types.DockerConfig{}, err
	}
	return types.DockerConfig{
		Image:         fmt.Sprintf("%s:%s", g.config.Name, g.config.Version),
		ContainerName: g.GetContainerName(),
		Ports:         g.config.Config.Ports,
		Volumes:       renderedVolumes,
		Environment:   g.config.Config.DependencyEnvironment,
	}, nil
}

// GetOnlineText returns the online text
func (g *genericDependency) GetOnlineText() string {
	return g.config.Config.OnlineText
}

// GetServiceEnvVariables returns the environment variables that need to
// be passed to services that use it
func (g *genericDependency) GetServiceEnvVariables() map[string]string {
	return g.config.Config.ServiceEnvironment
}
