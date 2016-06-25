package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) Revert(targetPaths []string, flags []string) error {
	commandArgs := []string{"svn", "revert"}
	for _, target := range targetPaths {
		commandArgs = append(commandArgs, target)
	}
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
