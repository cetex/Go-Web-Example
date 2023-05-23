# Go-Web-Example

## docs for the files used

### Dockerfile reference

https://docs.docker.com/engine/reference/builder/

### go.mod / go.sum

Info about go.mod / go.sum (used to version libary dependencies)

https://golangbyexample.com/go-mod-sum-module/

To create new go.mod:
`go mod init Go-Web-Example`
To update it:
`go mod tidy`

## Building docker image

`docker build . -t my-docker-image`

Explanation of options used:

```
  build       Build an image from a Dockerfile
  ...
  Usage:  docker build [OPTIONS] PATH | URL | -
  
  "." == PATH, Docker will reference "." (current directory) in build-context and will look for a Dockerfile in that directory.
  
  ...
    -t, --tag list                Name and optionally a tag in the 'name:tag' format
```

## Running docker image

`docker run -it -p 8080:8080 --rm my-docker-image`

Explanation of options used:


`build` - Build the dockerifle

```

  -i, --interactive                    Keep STDIN open even if not attached
  ...
  -p, --publish list                   Publish a container's port(s) to the host
  ...
  -t, --tty                            Allocate a pseudo-TTY
  ...
      --rm                             Automatically remove the container when it exits

```

Curl example:
```
$ curl localhost:8080/hello -v
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hello HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.85.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 23 May 2023 12:34:32 GMT
< Content-Length: 6
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello!
```

Post example posting two form strings named "test" and "address"
```
$ curl localhost:8080/form -v -d 'name=test&address=someaddress'
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /form HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.85.0
> Accept: */*
> Content-Length: 29
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 23 May 2023 13:00:43 GMT
< Content-Length: 58
< Content-Type: text/plain; charset=utf-8
<
POST request successful
Name = test
Address = someaddress
* Connection #0 to host localhost left intact
```

Post exmaple using json strings named "test" and "address"
```
$ curl localhost:8080/json -v -d '{"name": "test", "address": "someaddress"}'
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /json HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.85.0
> Accept: */*
> Content-Length: 42
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 23 May 2023 13:18:24 GMT
< Content-Length: 63
< Content-Type: text/plain; charset=utf-8
<
JSON POST request successful
Name = test
Address = someaddress
* Connection #0 to host localhost left intact
```


Post example using json string named "address" and value of "google.com"
```
$ curl localhost:8081/dns -v -d '{"domain": "google.com"}'
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081 (#0)
> POST /dns HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/7.85.0
> Accept: */*
> Content-Length: 24
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Tue, 23 May 2023 13:34:18 GMT
< Content-Length: 118
< Content-Type: text/plain; charset=utf-8
<
JSON POST request successful
Domain = google.com
[2a00:1450:400e:811::200e 142.250.179.142], 2a00:1450:400e:811::200e
* Connection #0 to host localhost left intact
```
