package utils

import "os"

type fileType struct {}

var File fileType

/*
Description: return true if path exist

 * Author: architect.bian
 * Date: 2018/08/06 16:16
 */
func (f fileType) Exist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
