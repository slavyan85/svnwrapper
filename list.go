package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

import (
	"strings"
)

func (svn *Svn) List(sourceUrl string, flags []string) ([]string, error) {
	var result []string
	commandArgs := []string{"svn", "list"}
	commandArgs = append(commandArgs, sourceUrl)
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	cmdResults, err := commandResult(commandArgs)
	if err != nil {
		return result, err
	}
	result = strings.Split(string(cmdResults), "\n")
	return result, nil
}
