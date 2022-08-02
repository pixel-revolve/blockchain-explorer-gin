package service

import (
	"gin/service/platform"
	"gin/service/system"
)

type Group struct {
	SystemServiceGroup   system.ServiceGroup
	PlatformServiceGroup platform.ServiceGroup
}

var GroupApp = new(Group)
