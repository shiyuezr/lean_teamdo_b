package routers

import (
	"github.com/kfchen81/beego/vanilla"
	"teamdo/rest/dev"
)

func init() {
	vanilla.Router(&dev.BDDReset{})
}
