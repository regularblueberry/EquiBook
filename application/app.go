package application

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfig struct {
	username string
}
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

type AppError struct {
	System AppSystemID;
	Code error;
	Message string;
}

func (self *App) ParseConfig(path string) {

	self.Config = AppConfig{
		username: "User",		
	};

	config, err := os.ReadFile(path);
	if err != nil{ 
		self.Logger.Println("Could not read config file using default config::", err)
		return 
	}

	err = json.Unmarshal(config, self)
	if err != nil{
		self.Logger.Println("Failed to parse json config file::", err)
		return 
	}
	
}


func (self *App) TryRecover(err AppError){

	switch err.System {
	case SYS_RENDER:
	case SYS_NETWORK:
	case SYS_STORAGE:
	default:
	}


}



