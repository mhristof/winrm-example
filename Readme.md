## success

```
$ docker build -f Dockerfile.success -t winrm.good .
$ docker run -it winrm.good  /go/src/foobar/foobar <windows ip> 'windows-password'
```

## failure

```
docker build -t winrm.bad .
docker run -it winrm.bad /go/src/foobar/foobar <windows ip> 'windows-password'
```
