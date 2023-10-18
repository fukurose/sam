# sam
sam is a file transfer tool.

![](./demo.gif)

## Installation
```
go install github.com/fukurose/sam@latest
```

## Getting Started

### host to which the file will be sent
Execute the following command.
```
sam received
```

### host that will receive the file
Execute the following command.
(default port is 50000)
```
sam bring -a 192.168.1.146:50000
```

## License
MIT
