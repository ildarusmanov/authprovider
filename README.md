# Authprovider

[![Build Status](https://travis-ci.org/ildarusmanov/authprovider.svg?branch=master)](https://travis-ci.org/ildarusmanov/authprovider)
[![Maintainability](https://api.codeclimate.com/v1/badges/a10ad1286a592257b2b1/maintainability)](https://codeclimate.com/github/ildarusmanov/authprovider/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a10ad1286a592257b2b1/test_coverage)](https://codeclimate.com/github/ildarusmanov/authprovider/test_coverage)

Token storage service

# Install
```
git clone https://github.com/ildarusmanov/authprovider
// or
go get github.com/ildarusmanov/authprovider
// move to project directory
cd [authprovider dicrectory]
// https://github.com/golang/dep
dep ensure
// run tests
go test -v ./grpcserver/ ./models/ ./providers ./services/
```

# Run with Docker

```
// move to directory
cd [authprovider directory]
// build
sudo docker build -t authprovider .
// run
sudo docker run -p 8000:8000 --network host authprovider 
```

