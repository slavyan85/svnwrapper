package svnwrapper

/*
author Kryuchenko Vyacheslav
*/

import (
	"encoding/xml"
)

type SvnStatusEntry struct {
	Path     string `xml:"path,attr"`
	WCStatus struct {
		Status string `xml:"item,attr"`
		//Revision string `xml:"revision,attr"`
	} `xml:"wc-status"`
}

type SvnStatusTarget struct {
	Path    string           `xml:"path,attr"`
	Entries []SvnStatusEntry `xml:"entry"`
}

type SvnStatus struct {
	XMLName xml.Name        `xml:"status"`
	Target  SvnStatusTarget `xml:"target"`
}

func (svn *Svn) Status(targetPath string, flags []string) (SvnStatus, error) {
	result := SvnStatus{}
	commandArgs := []string{"svn", "status"}
	commandArgs = append(commandArgs, targetPath)
	if len(flags) > 0 {
		for _, flag := range flags {
			commandArgs = append(commandArgs, flag)
		}
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
