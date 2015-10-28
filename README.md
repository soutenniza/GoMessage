# gomsg
Send logs from a file to an ip and port.

Download
```
go get github.com/soutenniza/gomsg
```

Usage
```
gomsg send <file> <ip:port>
```

Example for Cloud Foundry Syslogs
```
gomsg send syslog.txt 192.168.13.37:7000
```

Example for Graphite Events
```
gomsg send graphite.txt 192.168.13.37:7001
```
