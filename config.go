package repoman

/* Copyright (C) 2020 Pablo Alvarez de Sotomayor Posadillo

   This file is part of repoman.

   repoman is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   repoman is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with repoman. If not, see <http://www.gnu.org/licenses/>. */

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Repo struct {
	Name string `yaml:"name"`
	Repo string `yaml:"repository"`
}

type Group map[string]*Repo

type Config struct {
	Projects map[string]*Repo `yaml:"projects"`
	Groups   map[string]Group `yaml:"groups"`
}

var config Config

func parseConfig() error {
	if len(config.Groups) > 0 || len(config.Projects) > 0 {
		return nil
	}

	return parseConfigAt("./config.yml")
}

func parseConfigAt(file string) error {
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

	return err
}
