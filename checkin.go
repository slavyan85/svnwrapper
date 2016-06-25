package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) CheckIn(targetPaths []string, message string, flags []string) error {
	commandArgs := []string{"svn", "commit"}
	for _, target := range targetPaths {
		commandArgs = append(commandArgs, target)
	}
	commandArgs = append(commandArgs, "-m")
	commandArgs = append(commandArgs, message)
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
