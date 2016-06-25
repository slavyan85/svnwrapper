package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

import (
	"strings"
)

func (svn *Svn) Export(sourceUrl string, targetPath string, revision string, flags []string) error {
	commandArgs := []string{"svn", "export"}
	commandArgs = append(commandArgs, sourceUrl)
	commandArgs = append(commandArgs, targetPath)
	commandArgs = append(commandArgs, "-r")
	if strings.ToLower(revision) == "head" {
		commandArgs = append(commandArgs, "HEAD")
	} else {
		commandArgs = append(commandArgs, revision)
	}
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
