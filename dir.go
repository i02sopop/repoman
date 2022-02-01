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
	"os"
	"path"
	"strings"
)

func projectPath(group, name string) (string, error) {
	gpath := groupNameToPath(group)
	p := path.Join(gpath, name)
	if _, err := os.Stat(p); err != nil {
		return "", err
	}

	return p, nil
}

func mkGroupDir(group string) error {
	return mkDir(groupNameToPath(group))
}

func mkDir(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return err
	}

	fmt.Printf("Creating group: %s\n", path)

	return os.MkdirAll(path, os.ModePerm)
}

func groupNameToPath(group string) string {
	p := strings.Split(group, ".")
	return path.Join(p...)
}
