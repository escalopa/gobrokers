# gobrokers

A practice area of using golang with messsage brokers


## RabbitMQ Exampel 🐰

In this example we creaete a `group` to which many `consumers` can subscribe to. The `consumers` will receive messages from the `group` and process them.

So the message that is being published is received by all the `consumers` that are subscribed to the `group`, Which is the case in a group chat

### Setup 🔨 

1. Run rabbitmq server with docker compose

```bash
docker compose -f rabbitMQ/docker-compose.yml up
```

### Run 🚀

1. Run the publisher 📦

```bash
go run rabbitMQ/prod/main.go
```

2. Run the consumer 🍹

You can run multiple consumers by changing the `queueName` in the cli arg, Currently the `queueName` is set to `my-queue0` & `my-queue1`, But u can make more edits to the `rabbitMQ/prod/main.go` file to create more queues

```bash
go run rabbitMQ/consumer/main.go -queueName="my-queue0"
```