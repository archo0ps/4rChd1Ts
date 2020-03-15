package main

import (
	"4rChd1Ts/Scanf"
	"4rChd1Ts/config"
	"4rChd1Ts/help"
	"4rChd1Ts/utils"
	"flag"
	"github.com/asaskevich/govalidator"
	"github.com/fatih/color"
)

const (
	_RCHD1TS1 = `|	 _  _         ____ _         _ _ _____		|`
	_RCHD1TS2 = `|	| || |  _ __ / ___| |__   __| / |_   _|__	|`
	_RCHD1TS3 = "|	| || |_| '__| |   | '_ \\ / _` | | | |/ __|	|"
	_RCHD1TS4 = `|	|__   _| |  | |___| | | | (_| | | | |\__ \	|`
	_RCHD1TS5 = `|	   |_| |_|   \____|_| |_|\__,_|_| |_||___/	|`
	LINE      = `---------------------------------------------------------`
)

func main() {

	flag.Parse()
	if help.V {
		help.Version()
		return
	}
	if help.D == "" || help.U == "" || help.H {
		flag.Usage()
		return
	}
	config.RandomUserAgent = help.R
	if !govalidator.IsURL(help.U) {
		color.Red("Error:%s is not a url.", help.U)
		return
	}
	_, err := utils.DicFileToSlice(help.D, help.U)
	if err != nil {
		color.Red("Error:%s not found.", help.D)
		return
	}
	if help.U[:7] != "http://" {
		help.U = "http://" + help.U
	}

	color.Cyan(LINE)
	color.Cyan(_RCHD1TS1)
	color.Cyan(_RCHD1TS2)
	color.Cyan(_RCHD1TS3)
	color.Cyan(_RCHD1TS4)
	color.Cyan(_RCHD1TS5)
	color.Red("|					   Version:" + config.VERSION + "	|")
	color.Cyan(LINE)

	err = scanf.StartScanf(help.U, help.D)
	if err != nil {
		color.Red("Create %s file Error!", help.O)
	}

}
