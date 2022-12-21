package main

import (
	"flag"

	"github.com/srajeshinfotech/webui-golang/src/core"
	"github.com/srajeshinfotech/webui-golang/src/util"
)

var s_path *string

func parseCmdLineOptions(l_level *int) {

	flag.IntVar(l_level, "loglevel", 4, "Specify log level, 0-Panic, 1-Fatal, 2-Error, 3-Warn, 4-Info, 5-Debug, 6-Trace\n")
	s_path = flag.String("static", "./static", "Specify static pages path\n")
	flag.Parse()
}

func main() {
	var l_level int

	parseCmdLineOptions(&l_level)
	util.InitLogger(l_level)

	/* Start a webserver */
	core.StartWebServer(*s_path)
}
