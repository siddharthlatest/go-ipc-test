### Test Go / Node IPC concurrency

#### Pre-Requisites

Go 1.12, Node 10.16

#### Setup

`npm install lodash`

You might want to set file descriptor limit with `ulimit -n 1000` (depending on how you're scaling it)

#### Run

`go run main.go`

