# JIRA Time Manager CLI

## Installation

Installation is done by using the `go install` command and rename installed binary (from `jira-time-manager-cli` to `jtm`) in $GOPATH/bin:
```shell
go install github.com/ivanmagdic/jira-time-manager-cli
```

## Usage

### Initialization

```shell
jtm init --url <public-jira-url> -u <username> -p <password>
```

### List issues

```shell
jtm list-issues
```

### Start progress on issue

```shell
jtm start <issue-id>
```

### Stop progress on currently active issue
```shell
jtm stop
```