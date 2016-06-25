package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) MakeDir(targetUrl string, message string, flags []string) error {
	commandArgs := []string{"svn", "mkdir"}
	commandArgs = append(commandArgs, targetUrl)
	commandArgs = append(commandArgs, "-m")
	commandArgs = append(commandArgs, message)
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
