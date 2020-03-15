package help

import (
	"4rChd1Ts/config"
	"flag"
	"github.com/fatih/color"
)

var (
	H bool
	V bool
	U string
	D string
	O string
	R bool
)

func init(){
	flag.BoolVar(&H,"h",false,"this help")
	flag.BoolVar(&V, "v", false, "show version and exit")
	flag.BoolVar(&R, "r", false, "use random user-agent")

	flag.StringVar(&U,"u","","target url")
	flag.StringVar(&D,"d","","path to the dictionary `file`")
	flag.StringVar(&O,"o","","output `file` to write results to")
	flag.Usage = Usage
}

func Usage(){
	color.White("4rChd1Ts version:%s\nUsage:./4rChd1Ts [-u Url] " +
		"[-t Threads] [-d DictionaryFile] [-o OutputFile] [-r RandomUserAgent]\nOptions:",config.VERSION)
	flag.PrintDefaults()
}

func Version(){
	color.White("4rChd1Ts version:%s",config.VERSION)
}