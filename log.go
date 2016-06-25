package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

import (
	"encoding/xml"
)

type SvnLogEntry struct {
	Revision string       `xml:"revision,attr"`
	Author   string       `xml:"author"`
	Date     string       `xml:"date"`
	Message  string       `xml:"msg"`
	Paths    []SvnLogPath `xml:"paths>path"`
}

type SvnLogPath struct {
	//XMLName xml.Name `xml:"path"`
	Path   string `xml:",chardata"`
	Action string `xml:"action,attr"`
	Kind   string `xml:"kind,attr"`
}

type SvnLog struct {
	XMLName    xml.Name      `xml:"log"`
	LogEntries []SvnLogEntry `xml:"logentry"`
}

func (svn *Svn) Log(sourceUrl string, flags []string) (SvnLog, error) {
	result := SvnLog{}
	commandArgs := []string{"svn", "log"}
	commandArgs = append(commandArgs, sourceUrl)
	for _, flag := range flags {
		commandArgs = append(commandArgs, flag)
	}
	commandArgs = append(commandArgs, "--xml")
	xmlData, err := commandResult(commandArgs)
	if err != nil {
		return result, err
	}
	err = xml.Unmarshal(xmlData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
