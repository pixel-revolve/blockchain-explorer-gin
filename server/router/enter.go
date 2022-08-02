package router

import (
	"gin/router/platform"
	"gin/router/system"
)

type Group struct {
	System   system.RouterGroup
	Platform platform.RouterGroup
}

var GroupApp = new(Group)
