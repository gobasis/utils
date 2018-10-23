package utils

import (
	"strings"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"bytes"
)

/*
Description:
templateType
 * Author: architect.bian
 * Date: 2018/10/23 14:08
 */
type templateType struct {
	FuncMap template.FuncMap
}

/*
Description:
Template an instance of templateType
 * Author: architect.bian
 * Date: 2018/10/23 14:08
 */
var Template = templateType{
	FuncMap: template.FuncMap{
		"title": strings.Title,
		"contains": contains,
		"equal": equal,
	},
}

/*
Description:
ParseFiles parse a file, and apply funcMap return a new template
 * Author: architect.bian
 * Date: 2018/10/23 14:08
 */
func (t templateType) ParseFiles(file string) *template.Template {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	s := string(buf)
	name := filepath.Base(file)
	return template.Must(template.New(name).Funcs(t.FuncMap).Parse(s))
}

/*
Description:
parse file as template,return a result of string type
 * Author: architect.bian
 * Date: 2018/10/23 14:10
 */
func (t templateType) Execute(file string, data interface{}) string {
	var buf bytes.Buffer
	err := t.ParseFiles(file).Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

/*
Description:
whether s contain flag
 * Author: architect.bian
 * Date: 2018/08/06 16:18
 */
func contains(s interface{}, flag string) bool {
	if s == nil {
		return false
	}
	return strings.Contains(s.(string), flag)
}

/*
Description:
whether s equal with target
 * Author: architect.bian
 * Date: 2018/08/06 16:19
 */
func equal(s interface{}, target string) bool {
	if s == nil {
		return false
	}
	return s == target
}
