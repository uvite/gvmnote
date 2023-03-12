package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/adsafasdf"
	"github.com/flipped-aurora/gin-vue-admin/server/service/bots"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/other"
	"github.com/flipped-aurora/gin-vue-admin/server/service/strages"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	OtherServiceGroup     other.ServiceGroup
	BotsServiceGroup      bots.ServiceGroup
	StragesServiceGroup   strages.ServiceGroup
	AdsafasdfServiceGroup adsafasdf.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
