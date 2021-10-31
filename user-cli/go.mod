module github.com/869413421/laracom/user-cli

go 1.16

replace github.com/869413421/laracom/service => /E/go/src/github/869413421/laracom-master/laracom/user-service

require (
	github.com/869413421/laracom/user-service v0.0.0-20211031124734-9a97385df089
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
)
