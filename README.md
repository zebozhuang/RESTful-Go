#### 1. It is a server providing restful api
#### 2. Tt is a demo.
#### 3. Support

- [x] Listen http port
- [x] Listen unix sock
- [x] handle restful api
- [x] handle non-restful api

#### 4.usage

```
    go run server.go
```
```
    curl http://127.0.0.1/api/foo
```
```
    curl -XPOST http://127.0.0.1/api/foo
```

#### 5.Todo
```
1.wrap request and reponse as context
2.support middleware
3.support render web page 
4...
...
```

