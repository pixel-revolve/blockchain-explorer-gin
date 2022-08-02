package api

import "gin/api/system"

type Group struct {
	SystemApiGroup system.ApiGroup
}

var GroupApp = new(Group)
