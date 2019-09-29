package main

import (
	cf "github.com/cloudfoundry-community/go-cfenv"
	"github.com/spf13/viper"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)
func init() {
	log.SetOutput(os.Stdout)
	vcap,_ :=cf.Current()
	if cf.IsRunningOnCF(){
		initCFVcap(vcap)
	}else {
		viper.SetConfigFile(`local.json`)
		_ = viper.ReadInConfig()
	}
}
func main()  {
	// Hello, the web server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		str := fmt.Sprintf("Hello, %s !",viper.Get("test-cups.credentials.key"))
		io.WriteString(w, str)
	}
	http.HandleFunc("/", helloHandler)
	log.Panic(http.ListenAndServe(":8080", nil))
}
func initCFVcap(vcap *cf.App)  {
	//fetch all bound services
	vcapServices,_ := vcap.Services.WithNameUsingPattern(".*")
	for _, vcapService := range vcapServices {
		log.Println("Init vcap Service: ",vcapService.Name)
		//set vcap service to viper
		viper.Set(fmt.Sprintf("%s.credentials",vcapService.Name),vcapService)
	}
}