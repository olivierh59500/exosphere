package application

import (
	"fmt"
	"io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v2"

	"github.com/Originate/exosphere/src/config"
	"github.com/Originate/exosphere/src/docker/compose"
	"github.com/Originate/exosphere/src/docker/composebuilder"
	"github.com/Originate/exosphere/src/types"
	"github.com/Originate/exosphere/src/util"
)

// Initializer sets up the app
type Initializer struct {
	AppConfig                types.AppConfig
	BuiltAppDependencies     map[string]config.AppDependency
	DockerComposeConfig      types.DockerCompose
	DockerComposeProjectName string
	ServiceData              map[string]types.ServiceData
	ServiceConfigs           map[string]types.ServiceConfig
	AppDir                   string
	HomeDir                  string
	logChannel               chan string
}

// NewInitializer is Initializer's constructor
func NewInitializer(appConfig types.AppConfig, logChannel chan string, logRole, appDir, homeDir, dockerComposeProjectName string) (*Initializer, error) {
	serviceConfigs, err := config.GetServiceConfigs(appDir, appConfig)
	if err != nil {
		return &Initializer{}, err
	}
	appSetup := &Initializer{
		AppConfig:                appConfig,
		BuiltAppDependencies:     config.GetAppBuiltDependencies(appConfig, appDir, homeDir),
		DockerComposeConfig:      types.DockerCompose{Version: "3"},
		DockerComposeProjectName: dockerComposeProjectName,
		ServiceData:              appConfig.GetServiceData(),
		ServiceConfigs:           serviceConfigs,
		AppDir:                   appDir,
		HomeDir:                  homeDir,
		logChannel:               logChannel,
	}
	return appSetup, nil
}

func (i *Initializer) getAppDependenciesDockerConfigs() (types.DockerConfigs, error) {
	result := types.DockerConfigs{}
	for _, builtDependency := range i.BuiltAppDependencies {
		dockerConfig, err := builtDependency.GetDockerConfig()
		if err != nil {
			return result, err
		}
		result[builtDependency.GetContainerName()] = dockerConfig
	}
	return result, nil
}

// GetDockerConfigs returns the docker configs of all services and dependencies in
// the application
func (i *Initializer) GetDockerConfigs() (types.DockerConfigs, error) {
	dependencyDockerConfigs, err := i.getAppDependenciesDockerConfigs()
	if err != nil {
		return nil, err
	}
	serviceDockerConfigs, err := i.getServiceDockerConfigs()
	if err != nil {
		return nil, err
	}
	return dependencyDockerConfigs.Merge(serviceDockerConfigs), nil
}

func (i *Initializer) getServiceDockerConfigs() (types.DockerConfigs, error) {
	result := types.DockerConfigs{}
	for serviceName, serviceConfig := range i.ServiceConfigs {
		dockerComposeBuilder := composebuilder.NewDockerComposeBuilder(i.AppConfig, serviceConfig, i.ServiceData[serviceName], serviceName, i.AppDir, i.HomeDir)
		dockerConfig, err := dockerComposeBuilder.GetServiceDockerConfigs()
		if err != nil {
			return result, err
		}
		result = result.Merge(result, dockerConfig)
	}
	return result, nil
}

func (i *Initializer) renderDockerCompose(dockerComposeDir string) error {
	bytes, err := yaml.Marshal(i.DockerComposeConfig)
	if err != nil {
		return err
	}
	if err := util.CreateEmptyDirectory(dockerComposeDir); err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(dockerComposeDir, "docker-compose.yml"), bytes, 0777)
}

func (i *Initializer) setupDockerImages(dockerComposeDir string) error {
	opts := compose.BaseOptions{
		DockerComposeDir: dockerComposeDir,
		LogChannel:       i.logChannel,
		Env:              []string{fmt.Sprintf("COMPOSE_PROJECT_NAME=%s", i.DockerComposeProjectName)},
	}
	err := compose.PullAllImages(opts)
	if err != nil {
		return err
	}
	return compose.BuildAllImages(opts)
}

// Initialize sets up the entire app and returns an error if any
func (i *Initializer) Initialize() error {
	dockerConfigs, err := i.GetDockerConfigs()
	if err != nil {
		return err
	}
	i.DockerComposeConfig.Services = dockerConfigs
	dockerComposeDir := path.Join(i.AppDir, "tmp")
	if err := i.renderDockerCompose(dockerComposeDir); err != nil {
		return err
	}
	return i.setupDockerImages(dockerComposeDir)
}
