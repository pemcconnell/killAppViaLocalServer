# killchrome

This app will kill running chrome processes by the local URL `http://127.0.0.1:6732/chrome` being requested.

Tested in Ubuntu 14:04 & Windows 8.1

Three binarys are included in this repo: `./bin/killchrome` for Linux, `./bin/killchrome-32.exe` and `./bin/killchrome-64.exe` for Windows.

## Building Examples

####Linux

```
cd killchrome/
go build -o ../bin/killchrome
```

####Windows 32bit

```
cd killchrome/
GOOS=windows GOARCH=386 go build -o ../bin/killchrome-32.exe
```

####Windows 64bit

```
cd killchrome/
GOOS=windows GOARCH=386 go build -o ../bin/killchrome-64.exe
```

##Expected application Output

```
pete@home:~/Apps/killchrome$ ./bin/killchrome
2014/11/19 14:42:37 Listening on 127.0.0.1:6732
2014/11/19 14:42:39 killing chrome processes
2014/11/19 14:42:39 killed chrome processes
```
