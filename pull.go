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
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Pull clones and pulls all the repositories.
func Pull() error {
	mg.Deps(parseConfig)

	// Projects
	mg.Deps(cloneProjectRepos)
	mg.Deps(pullProjectRepos)

	// Groups
	mg.Deps(cloneGroupRepos)
	mg.Deps(pullGroupRepos)

	return nil
}

func pullGroupRepos() error {
	for group, repos := range config.Groups {
		for _, repo := range repos {
			if err := pullRepo(group, repo); err != nil {
				fmt.Printf("%s\n", err)
			}
		}
	}

	return nil
}

func pullProjectRepos() error {
	for _, repo := range config.Projects {
		if err := pullRepo(".", repo); err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	return nil
}

func pullRepo(group string, r *Repo) error {
	p, err := projectPath(group, r.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Pulling Repo: %s\n", p)

	return sh.Run("git", "-C", p, "pull")
}
