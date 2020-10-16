package repoman

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Repo struct {
	Name string `yaml:"name"`
	Repo string `yaml:"repository"`
}

type Group map[string]*Repo

type Config struct {
	Gopath      string           `yaml:"gopath"`
	StatusDepth int              `yaml:"status_depth"`
	Projects    map[string]*Repo `yaml:"projects"`
	Groups      map[string]Group `yaml:"groups"`
	pwd         string           `yaml:"-"` // current working, for reference
}

var config Config

func parseConfig() error {
	if len(config.Groups) > 0 || len(config.Projects) > 0 {
		return nil
	}

	return parseConfigAt("./config.yml")
}

func parseConfigAt(file string) error {
	config.Gopath = os.Getenv("GOPATH")
	if config.Gopath == "" {
		return errors.New("doesn't look like your GOPATH is configured")
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		return err
	}
	// Copy project and group Repo names
	for name, repo := range config.Projects {
		if repo.Name == "" {
			repo.Name = name
		}
	}
	for _, group := range config.Groups {
		for name, repo := range group {
			if repo.Name == "" {
				repo.Name = name
			}
		}
	}

	config.pwd, err = os.Getwd()
	return err
}
