# go-envdir

Envdir utility on Go

## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/go-envdir
 ```

## Usage

Usage example.

```
./go-envdir -c "sleep" -e "envList.txt"
```

## Available Methods

The following methods are available:

##### koind/go-envdir/exec

```go
Command(commandName, envDir string)
```