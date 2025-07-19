package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)



type AppConfig struct {
	username string
}


func (self *App) createConfig(path string){

	newfile, err := os.Create(path);
	if err != nil {
		self.Logger.Println("Could not create config file ::\n", err)
		return
	}
	jsondata, err := json.MarshalIndent(&self.Config, "", "");
	if err != nil {
		self.Logger.Println("Could not create config json data ::\n", err)
		return
	}
	byteswritten, err := newfile.Write(jsondata);
	if err != nil || byteswritten == 0{
		self.Logger.Println("Could not write out to config file ::\n", err)
		return
	}


}

func (self *App) parseConfig(path string){

	config, err := os.ReadFile(path);
	if err != nil{ 
		self.Logger.Println("Could not read config file ::\n", err)
		return 
	}

	err = json.Unmarshal(config, &self.Config)
	if err != nil{
		self.Logger.Println("Failed to parse json config file ::\n", err)
		return 
	}

}

func (self *App) initAppConfig(path string){

	self.Config = AppConfig{
		username: "User",		
	};
	_, err := os.Stat(path) 
	if err == nil {
		self.parseConfig(path + "/config.json")

	}else if errors.Is(err, os.ErrNotExist){
		fmt.Println("Config file was not found, Would you like to create one?");

	get_response:
		fmt.Print("[y/n] ");
		var response string;
		fmt.Scanln(&response)
		if(response != "y" && response != "n"){
			fmt.Println("Error! please reply with either a",
				    "'y' for yes or an 'n' for no",
			)
			goto get_response;
		}
		if(response == "y"){
			self.createConfig(path + "/config.json")
		}
	}else{
		self.Logger.Println("Error finding config file ::\n", err)
	}

}

func (self *App) LoadEnvironment(path string) {
	
	self.initAppConfig(path);

}
