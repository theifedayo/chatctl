### chatctl

## Description
chatctl allows you to chat with anyone via terminal.

## Installation

It can be installed by running

```
go install github.com/theifedayo/chatctl@latest 
```

## Usage
```bash
chatctl is a CLI for chat control -- server and client. This tool helps to send message across ips via command line interfaces - terminal, command prompt

Usage:
  chatctl [command]

Available Commands:
  client      join as client
  completion  Generate the autocompletion script for the specified shell
  generate    get your encrypted ip address
  help        Help about any command
  server      join as server

Flags:
  -h, --help   help for chatctl

Use "chatctl [command] --help" for more information about a command.
```
Terminal as server:
<img width="848" alt="Screenshot 2023-04-07 at 03 53 58" src="https://user-images.githubusercontent.com/47679952/230531757-7ba7b902-f191-4807-8c4d-0f72213293fa.png">


Terminal as client:
![WhatsApp Image 2023-04-07 at 03 54 14](https://user-images.githubusercontent.com/47679952/230531840-896a990e-456d-4159-8848-02db415fab55.jpeg)

## Contribution
You can go ahead to getIPAddress and encrypt with a key before sharing, then client decrypts the ip address with same key