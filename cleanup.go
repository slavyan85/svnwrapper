package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) Cleanup(targetPaths []string, flags []string) error {
	commandArgs := []string{"svn", "cleanup"}
	for _, target := range targetPaths {
		commandArgs = append(commandArgs, target)
	}
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
