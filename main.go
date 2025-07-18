package main

import (
	"log"
	"os"

	eq "github.com/regularblueberry/EquiBook/application"
	"github.com/regularblueberry/EquiBook/network"
	"github.com/regularblueberry/EquiBook/storage"
	"github.com/regularblueberry/EquiBook/ui"
)

func main(){

	var envPath string;
	var err eq.AppError;

	appdata := &eq.App{
		Logger: *log.New(os.Stdout, "EquiBook", log.LstdFlags),
	}

	if os.Args[0] == "" {
	    envPath, err.Code = os.Getwd();
	    if err.Code != nil {
		appdata.Logger.Println(err);
		return
	    }
	} else {
		envPath = os.Args[0]
	}

	appdata.ParseConfig(envPath + "/config.json")

	for {
		err = InitAppSystems(appdata)

		if err.Code != nil {
			appdata.TryRecover(err)
		}else{ break }
	}

	RunApp(appdata);
}

func InitAppSystems(app *eq.App) eq.AppError{
	err := eq.AppError{
		System:  eq.SYS_NULL,
		Code:    nil,
		Message: "No Error",
	}

	err = ui.InitRenderSystem(app)
	if err.Code != nil {
		return err;
	}
	err = storage.InitStorageSystem(app)
	if err.Code != nil {
		return err;
	}
	err = network.InitNetworkSystem(app)
	if err.Code != nil {
		return err;
	}

return err; 
}
