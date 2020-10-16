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
	"os"
	"path"

	"github.com/magefile/mage/sh"
)

func cloneGroupRepos() error {
	for group, repos := range config.Groups {
		if err := mkGroupDir(group); err != nil {
			return err
		}

		for _, repo := range repos {
			if err := cloneRepo(group, repo); err != nil {
				return err
			}
		}
	}

	return nil
}

func cloneProjectRepos() error {
	for _, repo := range config.Projects {
		if err := cloneRepo(".", repo); err != nil {
			return err
		}
	}

	return nil
}

func cloneRepo(group string, r *Repo) error {
	gpath := groupNameToPath(group)
	p := path.Join(gpath, r.Name)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		fmt.Printf("Cloning Repo %s into %s\n", r.Repo, p)

		return sh.Run("git", "-C", gpath, "clone", r.Repo, r.Name)
	}

	return nil
}
