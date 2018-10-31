package utils

import (
	"time"
	"math/rand"
	"runtime"
	_"net/http/pprof"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"github.com/gobasis/log"
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

/*
Description:
Get a channel that will be closed when a shutdown signal has been
triggered either from an OS signal such as SIGINT (Ctrl+C) or from
another subsystem such as the RPC server.
 * Author: architect.bian
 * Date: 2018/10/31 11:30
 */
func (_ runEnvType) NewInterruptChan() <-chan struct{} {
	c := make(chan struct{})
	go func() {
		// interruptSignals defines the default signals to catch in order to do a proper
		// shutdown.  This may be modified during init depending on the platform.
		var interruptSignals = []os.Signal{os.Interrupt}
		interruptChannel := make(chan os.Signal, 1)
		signal.Notify(interruptChannel, interruptSignals...)

		// Listen for initial shutdown signal and close the returned
		// channel to notify the caller.
		select {
		case sig := <-interruptChannel:
			log.Info("received signal. Now shutting down...", "signal", sig)
		}
		close(c)

		// Listen for repeated signals and display a message so the user
		// knows the shutdown is in progress and the process is not hung.
		for {
			select {
			case sig := <-interruptChannel:
				log.Info("received signal. Now shutting down...", "signal", sig)
			}
		}
	}()

	return c
}