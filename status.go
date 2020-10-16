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
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Status provides the status of each git repository.
func Status() error {
	mg.Deps(parseConfig)
	mg.Deps(statusProjectRepos)
	mg.Deps(statusGroupRepos)

	return nil
}

func statusGroupRepos() error {
	for group, repos := range config.Groups {
		for _, repo := range repos {
			if err := statusRepo(group, repo); err != nil {
				fmt.Printf("%s\n", err)
			}
		}
	}

	return nil
}

func statusProjectRepos() error {
	for _, repo := range config.Projects {
		if err := statusRepo(".", repo); err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	return nil
}

func statusRepo(group string, r *Repo) error {
	project, err := projectPath(group, r.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Status of %s\n", project)
	if err := sh.Run("git", "update-index", "-q", "--refresh"); err != nil {
		return err
	}

	remoteBranches, err := gitRemoteBranches(project)
	if err != nil {
		return err
	}

	branches, err := gitBranches(project)
	if err != nil {
		return err
	}

	for _, branch := range branches {
		pair := strings.Split(branch, " ")
		if _, ok := remoteBranches[pair[1]]; !ok {
			fmt.Printf("%s doesn't have an upstream\n", pair[0])
			continue
		}

		count, err := branchStatus(pair[0], pair[1])
		if err != nil {
			return err
		}

		fmt.Printf("%s is %s commits ahead and %s commits behind from %s\n",
			pair[0], count[0], count[1], pair[1])
	}

	return nil
}

func branchStatus(local, remote string) ([]string, error) {
	diff, err := sh.Output("git", "rev-list", "--left-right", "--count",
		"--branches="+local)
	if err != nil {
		return nil, err
	}

	return strings.Split(diff, "\t"), nil
}
