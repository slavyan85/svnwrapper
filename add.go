package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) Add(targetPaths []string, flags []string) error {
	commandArgs := []string{"svn", "add"}
	for _, target := range targetPaths {
		commandArgs = append(commandArgs, target)
	}
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
