#killAppViaLocalServer 

This project is designed to create a single binary which creates a local
webserver with an endpoint set up to kill a defined process. For the purposes
of this build this repo is set up to kill "chrome" and "chrome.exe" but this
can be updated in the `/killapp/main_<os>.go` files.

This instance will kill running chrome processes by the local URL `http://127.0.0.1:6732/kill` being requested.

Tested in Ubuntu 14:04 & Windows 8.1

Three binarys are included in this repo: `./bin/killapp` for Linux, `./bin/killapp-32.exe` and `./bin/killapp-64.exe` for Windows.

## Building Examples

####Linux

```
cd killapp/
go build -o ../bin/killapp
```

####Windows 32bit

```
cd killapp/
GOOS=windows GOARCH=386 go build -o ../bin/killapp-32.exe
```

####Windows 64bit

```
cd killapp/
GOOS=windows GOARCH=amd64 go build -o ../bin/killapp-64.exe
```

##Expected application Output

```
pete@home:~/Apps/killapp$ ./bin/killapp
2014/11/19 14:42:37 Listening on 127.0.0.1:6732
2014/11/19 14:42:39 killing chrome processes
2014/11/19 14:42:39 killed chrome processes
```
