#httptest

## About this service
This service makes http request to the url , hash response and prints hash value

## Used patterns
- Semaphone(to restrict number of workers)
- Decorator(Roundtrip)

#Running and building project

I have created a Make file to facilitate building and running the project

#Please use following command to get available options

```shell
make help
```

#Build Project

```shell
make build
```

## How to run service

```shell
go run main.go -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```
OR

```shell
./httptest -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```

```shell
./httptest adjust.com
```
