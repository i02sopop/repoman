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
	"strings"

	"github.com/magefile/mage/sh"
)

func gitBranches(project string) ([]string, error) {
	branches, err := sh.Output("git", "-C", project, "for-each-ref",
		`--format="%(refname:short) %(upstream:short)"`, "refs/heads")
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.ReplaceAll(branches, `"`, ""), "\n"), nil
}

func gitRemoteBranches(project string) (map[string]bool, error) {
	branches, err := sh.Output("git", "-C", project, "branch", "-r")
	if err != nil {
		return nil, err
	}

	remoteBranches := make(map[string]bool)
	for _, branch := range strings.Split(strings.ReplaceAll(branches, `"`, ""), "\n") {
		remoteBranches[strings.TrimSpace(branch)] = true
	}

	return remoteBranches, nil
}
