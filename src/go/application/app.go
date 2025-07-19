package application

import (
	"log"
)


type App struct {
	Logger log.Logger;
	Config AppConfig;
}

type AppSystemID int;
const (
	SYS_NULL = iota
	SYS_RENDER  
	SYS_NETWORK
	SYS_STORAGE
)



