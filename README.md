# killchrome

This app will kill running chrome processes by the local URL `http://127.0.0.1:6732/chrome` being requested.

Tested in Ubuntu 14:04 & Windows 8.1

Two binarys are included in this repo: `/bin/killchrome` for Linux and `/bin/killchrome.exe` for Windows.

###Example Output

```
pete@home:~/Apps/killchrome$ ./bin/killchrome
2014/11/19 14:42:37 Listening on 127.0.0.1:6732
2014/11/19 14:42:39 killing chrome processes
2014/11/19 14:42:39 killed chrome processes
```
