package main

import (
	cf "github.com/cloudfoundry-community/go-cfenv"
	"github.com/spf13/viper"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	log.Println(os.Getenv("ENV_TEST"))
}
func initCFVcap(vcap *cf.App)  {
	//fetch all bound services
	vcapServices,_ := vcap.Services.WithNameUsingPattern(".*")
	for _, vcapService := range vcapServices {
		log.Println("Init vcap Service: ",vcapService.Name)
		//set vcap service to viper
		viper.Set(fmt.Sprintf("%s.credentials",vcapService.Name),vcapService.Credentials)
	}
}
func main()  {
	// Hello, the web server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		str := fmt.Sprintf("Hello, cups: %s,env var: %s !",viper.Get("test-cups.credentials.key"),viper.Get("env.test"))
		io.WriteString(w, str)
	}
	http.HandleFunc("/", helloHandler)
	log.Panic(http.ListenAndServe(":8080", nil))
}
