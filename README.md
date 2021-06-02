# CIDR Range Calculator
ipcalc is a CIDR range calculator.

## Installation
You must have Go installed and $GOPATH/bin must be in your path. For example:
`export PATH=$PATH:$(go env GOPATH)/bin`

Download and install the `ipcalc` binary

`git clone git@github.com:andy-trimble/ipcalc.git`

`cd ipcalc`

`go install`

## Usage
```
> ipcalc 10.1.0.0/24
CIDR Range: 10.1.0.0/24
Netmask:    255.255.255.0
First IP:   10.1.0.0
Last IP:    10.1.0.255
Addresses:  256
```

```
> echo "10.1.0.0/24" | ipcalc
CIDR Range: 10.1.0.0/24
Netmask:    255.255.255.0
First IP:   10.1.0.0
Last IP:    10.1.0.255
Addresses:  256
```
