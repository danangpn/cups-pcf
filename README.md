
##Sample CUPS PCF

## Dependencies
- [Go 1.9](https://golang.org/dl/) or later.
- [go-cfenv](https://github.com/cloudfoundry-community/go-cfenv) : Go CF Environment
- [viper](https://github.com/spf13/viper) : Go configuration library.
- [dep](https://github.com/golang/dep) : dependencies management.

## Install (Manual/local)
1. ```go get github.com/danangpn/cups-pcf``` or clone this repository.
2. dep ensure -v
3. go build -o app
4. run the CLI : ./app


## Deploy in PCF
1. ```go get github.com/danangpn/cups-pcf``` or clone this repository.
2. dep ensure -v
3. cf cups test-cups -p '{"key":"value-of-test-cups-key"}'
4. cf push

