# go-rabbitmq-exam
[![Go](https://img.shields.io/badge/go-1.21-blue.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/dl/)
[![amqp091-go](https://img.shields.io/badge/amqp091go-0.9.2-orange.svg?style=for-the-badge&logo=go&logoColor=white)](https://github.com/rabbitmq/amqp091-go)

> #### 🎯 docker-compose, cobra 를 사용한 rabbitmq example 입니다.
> #### 🎯 go-rabbitmq-exam using docker-compose and cobra.


## ⚙️ Installation
```shell
$ git clone https://github.com/LimJiAn/go-rabbitmq-exam
```
## 👀 Usage
#### 1. Run RabbitMQ Server (docker-compose.yml)
```bash
$ docker compose build
```
```bash
$ docker compose up
```
#### 2. Wait 1-2 minutes
```console
[+] Running 1/0
 ✔ Container rabbitmq  Created                                                                                                                               0.0s
Attaching to rabbitmq
...
...
...
...
rabbitmq  |  completed with 4 plugins.
rabbitmq  | 2023-09-06 09:25:43.553462+00:00 [info] <0.474.0> Server startup complete; 4 plugins started.
rabbitmq  | 2023-09-06 09:25:43.553462+00:00 [info] <0.474.0>  * rabbitmq_prometheus
rabbitmq  | 2023-09-06 09:25:43.553462+00:00 [info] <0.474.0>  * rabbitmq_management
rabbitmq  | 2023-09-06 09:25:43.553462+00:00 [info] <0.474.0>  * rabbitmq_management_agent
rabbitmq  | 2023-09-06 09:25:43.553462+00:00 [info] <0.474.0>  * rabbitmq_web_dispatch

```
#### 3. You can see useful RabbitMQ dashboard at [localhost:15672](http://localhost:15672)
![Rabbitmq dashboard](https://github.com/LimJiAn/go-rabbitmq-exam/assets/85569173/73e8f14c-98e8-4c12-a59f-eca323a54816)

#### 4. Command
* ###### Run publish, consume
```
$ go run main.go --help
go rabbitmq exam is a CLI tool for rabbitmq exam.

Usage:
  go-rabbitmq-exam [flags]
  go-rabbitmq-exam [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  consume     receive message from queue
  help        Help about any command
  publish     Send message to queue.

Flags:
  -h, --help     help for go-rabbitmq-exam
  -t, --toggle   Help message for toggle

Use "go-rabbitmq-exam [command] --help" for more information about a command.
```
---
* ###### Run publish
```
$ go run main.go publish --help
publish.go, consume.go together make a simple example of using RabbitMQ

Usage:
  go-rabbitmq-exam publish [flags]

Flags:
  -c, --count int           count of message to send (default 1)
  -e, --exchange string     exchange name
  -h, --help                help for publish
  -q, --queue string        queue name (default "hello")
  -r, --routingkey string   routing key (default "info")
  -t, --type string         exchange type [direct, fanout, topic, headers] (default "direct")
```
---
* ###### Run consume
```
$ go run main.go consume --help
publish.go, consume.go together make a simple example of using RabbitMQ

Usage:
  go-rabbitmq-exam consume [flags]

Flags:
  -e, --exchange string     exchange name
  -h, --help                help for consume
  -q, --queue string        queue name (default "hello")
  -r, --routingkey string   routing key (default "info")
  -t, --type string         exchange type (default "direct")
```
#### 5. Example
* ###### Without flag
```bash
$ go run main.go publish
$ go run main.go consume
```
```console
2023/09/07 20:00:00  📧 Sent Hello World!!
2023/09/07 20:00:00  ✋ Waiting for messages. To exit press CTRL+C
2023/09/07 20:00:00  🆗 Received a message: Hello World!! / Count: 1
```
---
* ###### With flag
```bash
$ go run main.go publish -e=test -c=3
$ go run main.go consume -e=test
```
```console
2023/09/08 20:00:00  📮 exchangeName: test, routingKey: info, exchangeType: direct, count: 3
2023/09/08 20:00:00  📧 Sent Hello World!! [1]
2023/09/08 20:00:00  📧 Sent Hello World!! [2]
2023/09/08 20:00:00  📧 Sent Hello World!! [3]
2023/09/08 20:00:00  📮 exchangeName: test, routingKey: info, exchangeType: direct, queue: hello
2023/09/08 20:00:00  ✋ Waiting for messages. To exit press CTRL+C
2023/09/08 20:00:00  🆗 Received a message: Hello World!! [1] / Count: 1
2023/09/08 20:00:00  🆗 Received a message: Hello World!! [2] / Count: 2
2023/09/08 20:00:00  🆗 Received a message: Hello World!! [3] / Count: 3

```
## 📚 Reference
#### [RabbitMQ](https://www.rabbitmq.com/)
