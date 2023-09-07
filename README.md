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
![Rabbitmq dashboard](https://github.com/LimJiAn/go-rabbitmq-exam/assets/85569173/35ab9269-e1ae-4b52-8749-5ede37e2dae7)

#### 4. Command
```
$ go run main.go --help
go rabbitmq exam is a CLI tool for rabbitmq exam.

Usage:
  go-rabbitmq-exam [flags]
  go-rabbitmq-exam [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  consume     A brief description of your command
  help        Help about any command
  publish     asassasasA brief description of your command

Flags:
  -h, --help     help for go-rabbitmq-exam
  -t, --toggle   Help message for toggle

Use "go-rabbitmq-exam [command] --help" for more information about a command.
```
## 📚 Reference
#### [RabbitMQ](https://www.rabbitmq.com/)
