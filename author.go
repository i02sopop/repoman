package repoman

/* Copyright (C) 2020-2022 Pablo Alvarez de Sotomayor Posadillo

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

// ChangeAuthor change the repositories author configurations.
func ChangeAuthor() error {
	mg.Deps(parseConfig)

	return nil
}

func changeAuthorGroupRepos() error {
	for group, repos := range config.Groups {
		for _, repo := range repos {
			if err := changeAuthor(group, repo); err != nil {
				return err
			}
		}
	}

	return nil
}

func changeAuthorProjectRepos() error {
	for _, repo := range config.Projects {
		if err := changeAuthor(".", repo); err != nil {
			return err
		}
	}

	return nil
}

func changeAuthor(group string, r *Repo) error {
	p, err := projectPath(group, r.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Changing author for Repo %s\n", p)
	if err := sh.Run("git", "-C", p, "config", "user.email", config.Author.Email); err != nil {
		return err
	}

	return sh.Run("git", "-C", p, "config", "user.name", config.Author.Name)
}
