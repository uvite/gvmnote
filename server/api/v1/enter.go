package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/bots"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/other"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/vite"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	OtherApiGroup   other.ApiGroup
	BotsApiGroup    bots.ApiGroup
	ViteApiGroup    vite.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
