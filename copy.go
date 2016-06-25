package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

func (svn *Svn) Copy(sourceUrl string, targetUrl string, message string, flags []string) error {
	commandArgs := []string{"svn", "copy"}
	commandArgs = append(commandArgs, sourceUrl)
	commandArgs = append(commandArgs, targetUrl)
	commandArgs = append(commandArgs, "-m")
	commandArgs = append(commandArgs, message)
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	return runCommand(commandArgs)
}
