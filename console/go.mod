module console

go 1.13

replace (
	github.com/mix-go/bean => ../../mix/src/bean
	github.com/mix-go/console => ../../mix/src/console
	github.com/mix-go/dotenv => ../../mix/src/dotenv
	github.com/mix-go/event => ../../mix/src/event
	github.com/mix-go/logrus => ../../mix/src/logrus
	github.com/mix-go/workerpool => ../../mix/src/workerpool
)

require (
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/mix-go/bean v1.0.0-beta5
	github.com/mix-go/console v1.0.0-beta8
	github.com/mix-go/dotenv v0.0.0-00010101000000-000000000000
	github.com/mix-go/event v1.0.0-beta5
	github.com/mix-go/logrus v1.0.0-beta5
	github.com/mix-go/workerpool v1.0.0-beta9
)
