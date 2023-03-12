package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/adsafasdf"
	"github.com/flipped-aurora/gin-vue-admin/server/router/bots"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/other"
	"github.com/flipped-aurora/gin-vue-admin/server/router/strages"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Other     other.RouterGroup
	Bots      bots.RouterGroup
	Strages   strages.RouterGroup
	Adsafasdf adsafasdf.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
