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

## Contribution
You can go ahead to getIPAddress and encrypt with a key before sharing, then client decrypts the ip address with same key