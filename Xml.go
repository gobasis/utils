package utils

import (
	"strings"
)

type xmlType struct {}

var Xml xmlType

/*
Description: 返回参数s中标签tag之间内容

 * Author: architect.bian
 * Date: 2018/08/06 16:19
 */
func (x xmlType) TagContent(s string, tag string) string {
	var beginTag = "<" + tag
	var endTag = "</" + tag + ">"
	fromPos := strings.Index(s, beginTag)
	var beginTagEnd = ">"
	for ; fromPos < len(s); fromPos++ {
		if s[fromPos:fromPos +len(beginTagEnd)] == beginTagEnd {
			//fmt.Println(s[fromPos:fromPos +len(beginTagEnd)])
			fromPos++
			break
		}
	}
	toPos := strings.Index(s, endTag)
	if toPos != -1 {
		return s[fromPos:toPos]
	} else {
		return s[fromPos:]
	}
}

/*
Description: 抽取tag的属性及返回tag标签

 * Author: architect.bian
 * Date: 2018/08/06 16:20
 */
func (x xmlType) TagAtrrs(html string, tag string, pos int) (eleFragment string, attrFragment string, position int) {
	var beginTag = "<" + tag
	var endTag = ">"
	var posBegin = -1
	var posEnd = -1
	if len(html) < len(beginTag) + len(endTag) || len(html) == pos {
		return "", "", pos
	}
	for ; pos < len(html); pos++ {
		if pos < len(beginTag) {
			continue
		}
		from := pos - len(beginTag)
		if html[from:pos] == beginTag {
			posBegin = pos
		}
		if posBegin != -1 {
			to := pos + len(endTag)
			if html[pos:to] == endTag {
				posEnd = pos
			}
		}
		if posEnd != -1 {
			break
		}
	}
	if posBegin != -1 && posEnd != -1 {
		if posBegin == posEnd {
			return beginTag + endTag, "", pos
		}
		attrFragment := html[posBegin:posEnd]
		eleFragment := beginTag + attrFragment + endTag
		return eleFragment, attrFragment, pos
	} else {
		return "", "", pos
	}
}
