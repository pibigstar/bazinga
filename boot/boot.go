package boot

import (
	"github.com/gogf/gf/frame/g"
	"runtime"
)

func init() {
	setConfig()
	//startPProf()
}

func setConfig() {
	if runtime.GOOS == "linux" {
		g.Cfg().SetFileName("config/config-prod.toml")
	}else  {
		g.Cfg().SetFileName("config/config-dev.toml")
	}
}

func startPProf() {
	g.Server().EnablePProf()
}
