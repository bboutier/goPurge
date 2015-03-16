# goPurge
Simple golang script to do parallel purge, based on a configuration file listing the paths to purge with associate delay (in days). Parallel degree configurable.

###Build:

```
go get github.com/op/go-logging
go build purge-files.go
````

### Run:
```
./purge-files purge-conf.json
```

### Output:
```
2015-03-15 22:09:45 5938 INFO 001 Start purge
2015-03-15 22:09:45 5938 INFO 002 Waiting parallel purge to finish...
2015-03-15 22:09:45 5938 INFO 004 Process line: path=/tmp/test_purge/log/*, delay=10
2015-03-15 22:09:45 5938 INFO 003 Process line: path=/tmp/test_purge/*, delay=30
2015-03-15 22:09:45 5938 INFO 005 File /tmp/test_purge/log/test.log deleted successfully
2015-03-15 22:09:45 5938 INFO 006 Process line: path=/tmp/test_purge2/*.xml, delay=5
2015-03-15 22:09:45 5938 INFO 007 File /tmp/test_purge/toto1 deleted successfully
2015-03-15 22:09:45 5938 INFO 008 File /tmp/test_purge/toto2 deleted successfully
2015-03-15 22:09:45 5938 INFO 009 File /tmp/test_purge2/toto.xml deleted successfully
2015-03-15 22:09:45 5938 INFO 00a Purge done.
```

### Example of configuration file:

```json
{
  "parallel_degree": 2,
  "purge_info": [
    {
      "path": "/tmp/test_purge/*",
      "delay": 30
    },
    {
      "path": "/tmp/test_purge/log/*",
      "delay": 10
    },
    {
      "path": "/tmp/test_purge2/*.xml",
      "delay": 5
    }
  ]
}
```
