# Go Curl
Go solution to Coding Challenges [build your own curl](https://codingchallenges.fyi/challenges/challenge-curl)

## Testing
As this a complete solution the intermediate build steps aren't shown.

### Get
```bash
./gocurl http://eu.httpbin.org:80/get
{
  "args": {}, 
  "headers": {
    "Accept": "*/*", 
    "Accept-Encoding": "gzip", 
    "Host": "eu.httpbin.org", 
    "User-Agent": "Go-http-client/1.1", 
    "X-Amzn-Trace-Id": "Root=1-67dacedb-3b6aedeb399ac2230dc7e103"
  }, 
  "origin": "84.65.3.31", 
  "url": "http://eu.httpbin.org/get"
}
```

### Get With Verbose
### Get
```bash
./gocurl http://eu.httpbin.org:80/get
-v  http://eu.httpbin.org:80/get

> GET /get HTTP/1.1
> Host: eu.httpbin.org:80
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Content-Length: 296
< Server: gunicorn/19.9.0
< Access-Control-Allow-Origin: *
< Access-Control-Allow-Credentials: true
< Date: Wed, 19 Mar 2025 14:05:59 GMT
{
  "args": {}, 
  "headers": {
    "Accept": "*/*", 
    "Accept-Encoding": "gzip", 
    "Host": "eu.httpbin.org", 
    "User-Agent": "Go-http-client/1.1", 
    "X-Amzn-Trace-Id": "Root=1-67dacf45-7561a5812155b18e37eaa560"
  }, 
  "origin": "84.65.3.31", 
  "url": "http://eu.httpbin.org/get"
}
```

### Delete
```bash
./gocurl -X DELETE http://eu.httpbin.org/delete
{
  "args": {}, 
  "data": "", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "*/*", 
    "Accept-Encoding": "gzip", 
    "Host": "eu.httpbin.org", 
    "User-Agent": "Go-http-client/1.1", 
    "X-Amzn-Trace-Id": "Root=1-67dad006-70e6a1804d2c476956e60aa7"
  }, 
  "json": null, 
  "origin": "84.65.3.31", 
  "url": "http://eu.httpbin.org/delete"
}

```
### POST
```bash
./gocurl -X POST -d '{"key": "value"}' -H "Content-Type: application/json" http://eu.httpbin.org/post
{
  "args": {}, 
  "data": "{\"key\": \"value\"}", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "*/*", 
    "Accept-Encoding": "gzip", 
    "Content-Length": "16", 
    "Content-Type": "application/json", 
    "Host": "eu.httpbin.org", 
    "User-Agent": "Go-http-client/1.1", 
    "X-Amzn-Trace-Id": "Root=1-67dae005-205075f95f02b92375da67d3"
  }, 
  "json": {
    "key": "value"
  }, 
  "origin": "84.65.3.31", 
  "url": "http://eu.httpbin.org/post"
}
```


### PUT
```bash
./gocurl -X PUT -d '{"key2": "value2"}' -H "Content-Type: application/json" http://eu.httpbin.org/put
{
  "args": {}, 
  "data": "{\"key2\": \"value2\"}", 
  "files": {}, 
  "form": {}, 
  "headers": {
    "Accept": "*/*", 
    "Accept-Encoding": "gzip", 
    "Content-Length": "18", 
    "Content-Type": "application/json", 
    "Host": "eu.httpbin.org", 
    "User-Agent": "Go-http-client/1.1", 
    "X-Amzn-Trace-Id": "Root=1-67dafce4-7bd090145f1a7f09479776e2"
  }, 
  "json": {
    "key2": "value2"
  }, 
  "origin": "84.65.3.31", 
  "url": "http://eu.httpbin.org/put"
}
```





