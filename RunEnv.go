package utils

import (
	"time"
	"math/rand"
	"runtime"
	_"net/http/pprof"
	"fmt"
	"net/http"
)

type runEnvType struct {}

/*
Description:
A public instance of runEnvType
 * Author: architect.bian
 * Date: 2018/10/30 19:40
 */
var RunEnv runEnvType

/*
Description:
setup run environment
 * Author: architect.bian
 * Date: 2018/08/06 16:16
 */
func (_ runEnvType) Setup(kOrv ...interface{}) {
	rand.Seed(time.Now().UnixNano())
	// Setting a higher number here allows more disk I/O calls to be scheduled, hence considerably
	// improving throughput. The extra CPU overhead is almost negligible in comparison. reference:
	// https://groups.google.com/forum/#!topic/golang-nuts/jPb_h3TvlKE
	runtime.GOMAXPROCS(512)
	var params = make(map[string]interface{})
	for i, p := range kOrv {
		if i % 2 == 0 {
			params[fmt.Sprint(p)] = nil
			if len(kOrv) > i + 1 {
				params[fmt.Sprint(p)] = kOrv[i + 1]
			}
		}
	}
	if params["pprof"] != nil { //startup pprof
		go func() {
			http.ListenAndServe(fmt.Sprintf(":%v", params["pprof"]), nil)
		}()
	}
}