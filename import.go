package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) Import(targetUrl string, sourcePath string, message string, flags []string) error {
	commandArgs := []string{"svn", "import"}
	commandArgs = append(commandArgs, sourcePath)
	commandArgs = append(commandArgs, targetUrl)
	commandArgs = append(commandArgs, "-m")
	commandArgs = append(commandArgs, message)
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
