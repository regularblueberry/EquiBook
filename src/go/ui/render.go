package ui

import (
	eq "github.com/regularblueberry/EquiBook/application"
)

// #include "../../C/gfx.h"
import "C"

func InitRenderSystem(appdata *eq.App) eq.AppError{
	C.

return eq.AppError{
	System:  eq.SYS_RENDER, 
	Code: 	 nil, 
	Message: "No Error",
}
}
