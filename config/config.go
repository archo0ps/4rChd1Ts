package config

const (
	VERSION   = "0.1"
	DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36"
)

var RandomUserAgent = false
var DefaultMaxWorker = 1
var UserAgents = []string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3528.4 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.123 Safari/537.36 EdgA/42.0.0.2305",
	"Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Ubuntu/11.04 Chromium/10.0.648.205 Chrome/10.0.648.205 Safari/534.16",
	"Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0",
	"Opera/9.63 (X11; Linux x86_64; U; ru) Presto/2.1.1",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36 OPR/51.0.2830.55",
	"Mozilla/5.0 (Windows NT 7.1; Win64; x64) AppleWebKit/536.14 (KHTML, like Gecko) Chrome/35.0.2022.26 Safari/536.14",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 YaBrowser/19.12.2.252 Yowser/2.5 Safari/94959E",
	"Outlook-Express/7.0 (MSIE 7.0; Windows NT 6.1; Trident/7.0; SLCC2; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; .NET4.0C; .NET4.0E; Tablet PC 2.0; TmstmpExt)",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; MAGWJS; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 AOL/9.8 AOLBuild/4346.2019.US Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; WebView/3.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) CriOS/59.0.3071.102 Mobile/14F89 Safari/602.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_1_2 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 YaBrowser/18.4.1.542.10 Mobile/15B202 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 10; Pixel 3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9; Pixel 3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3836.3 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 10; Pixel 3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.186 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; U; Android 8.1.0; en-gb; Redmi Note 5 Pro Build/OPM1.171019.011) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.141 Mobile Safari/537.36 XiaoMi/MiuiBrowser/11.4.3-g",
}
var DefaultBackUpName = []string{".zip",".rar",".tar",".7z",".tar.xz",".tar.gz",".bak"}
