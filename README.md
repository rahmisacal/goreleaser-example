# gorestapi
A Simple REST API Build Using Golang

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

## About gorestapi:
- a simple music restapi that retrives info about the author, album name, price of it

## Author:
- rahmisacal [github](https://github.com/rahmisacal)


## Get response about the albums
- get response of the albums that are stored locally
```
$ curl http://localhost:1025/api/books
$ curl http://localhost:1025/api/books/1
$ curl http://localhost:1025/api/books/2
```

## Post response about the albums
- post response of the albums that are stored locally
```
$ curl http://localhost:1025/api/books --request POST --header "Content-Type: application/json" --data '{"id":"3","isbn":"398743","title":"Book three","author":{"firstname":"roe","lastname":"hoe"}}'
```

## Put response about the albums
- Put response of the albums that are stored locally
```
$ curl http://localhost:1025/api/books --request PUT --header "Content-Type: application/json" --data '{"id":"3","isbn":"398743","title":"Book three","author":{"firstname":"zoe","lastname":"hoe"}}'
```