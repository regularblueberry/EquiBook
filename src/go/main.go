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

	var envPath string = ".";
	var err eq.AppError;

	appdata := &eq.App{
		Logger: *log.New(os.Stdout, "EquiBook", log.LstdFlags),
	}

	if len(os.Args) > 1 && os.Args[1] != "" {
		envPath = os.Args[1]
	}

	appdata.LoadEnvironment(envPath)

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
