package svnwrapper

/*
   author Vyacheslav Kryuchenko
*/

import (
	"encoding/xml"
)

type SvnInfoEntryCommit struct {
	Revision string `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

type SvnInfoEntryRepo struct {
	Root string `xml:"root"`
	UUID string `xml:"uuid"`
}

type SvnInfoEntry struct {
	Revision    string             `xml:"revision,attr"`
	Path        string             `xml:"path,attr"`
	Kind        string             `xml:"kind,attr"`
	URL         string             `xml:"url"`
	RelativeURL string             `xml:"relative-url"`
	Repository  SvnInfoEntryRepo   `xml:"repository"`
	Commit      SvnInfoEntryCommit `xml:"commit"`
}

type SvnInfo struct {
	XMLName xml.Name `xml:"info"`
	Entry   struct {
		Revision    string             `xml:"revision,attr"`
		Path        string             `xml:"path,attr"`
		Kind        string             `xml:"kind,attr"`
		URL         string             `xml:"url"`
		RelativeURL string             `xml:"relative-url"`
		Repository  SvnInfoEntryRepo   `xml:"repository"`
		Commit      SvnInfoEntryCommit `xml:"commit"`
	} `xml:"entry"`
}

func (svn *Svn) Info(sourceUrl string, flags []string) (SvnInfo, error) {
	result := SvnInfo{}
	commandArgs := []string{"svn", "info"}
	commandArgs = append(commandArgs, sourceUrl)
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
