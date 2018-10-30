package utils

import (
	"time"
	"math/rand"
	"runtime"
)

type runEnvType struct {}

var RunEnv runEnvType

/*
Description:
setup run environment
 * Author: architect.bian
 * Date: 2018/08/06 16:16
 */
func (_ runEnvType) Setup() {
	rand.Seed(time.Now().UnixNano())
	// Setting a higher number here allows more disk I/O calls to be scheduled, hence considerably
	// improving throughput. The extra CPU overhead is almost negligible in comparison. reference:
	// https://groups.google.com/forum/#!topic/golang-nuts/jPb_h3TvlKE
	runtime.GOMAXPROCS(512)
}