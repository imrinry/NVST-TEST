# NVST-TEST
  Hello, I am glad you are interested in my exam results.

You can run the following command to test the program.
  ```sh
go test -v
```
You can run the following command to test the program's performance.
  ```sh
go test -bench=.
```


# Structure
```bash
├── handlers
│   ├── user.go
│   ├── interface.go
├── services
│   ├── user.go
│   ├── interface.go
├── repositories
│   ├── user.go
│   ├── garmin.go
│   ├── interface.go
├── models
│   ├── user.go
├── router
│   ├── router.go
├── middlewares
│   ├── middlewares.go
├── storage
│   ├── db.go
│   ├── redis.go
├── utils
│   ├── utils.go
├── main.go
├── go.mod
├── go.sum
├── .env
├── .gitlab-ci.yml
└── .gitignore
```


