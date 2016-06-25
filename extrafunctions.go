package svnwrapper

/*
author Kryuchenko Vyacheslav
*/

import (
	"os"
	"regexp"
)

func (svn *Svn) HardReset(targetPaths []string, deleteExternal bool) error {
	removeSignature := &regexp.Regexp{}
	if deleteExternal {
		rs, err := regexp.Compile(`(unversioned|ignored|external)`)
		if err != nil {
			return err
		}
		removeSignature = rs
	} else {
		rs, err := regexp.Compile(`(unversioned|ignored)`)
		if err != nil {
			return err
		}
		removeSignature = rs
	}
	err := svn.Cleanup(targetPaths, []string{})
	if err != nil {
		return err
	}
	err = svn.Revert(targetPaths, []string{"-R"})
	if err != nil {
		return err
	}
	for _, target := range targetPaths {
		status, err := svn.Status(target, []string{"--no-ignore"})
		if err != nil {
			return err
		}
		entiersList := status.Target.Entries
		if len(entiersList) > 0 {
			for _, item := range entiersList {
				if removeSignature.MatchString(item.WCStatus.Status) {
					err := os.RemoveAll(item.Path)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
