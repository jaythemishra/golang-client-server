# golang-client-server

This is a simple web server and client program written in Go. It allows the user to send API calls to the server to create jobs and assign them to specific instances of the client program by client id. The client then issues a GET request to the server to get its job if it has been assigned one, carries out all tasks, and sends back a summary of the results. Right now the supported task types are `ping` and `curl`, each of which can test the level 3 and level 7 network connectivity respectively. Included are the files required to package the client program in a linuxkit VM that can then be deployed and run on any device.

## Pre-requisites

- Make sure you have `go` installed on your machine
- Make sure you have `linuxkit`, `docker`, and `qemu` installed on your machine if you want to build the linuxkit package and VM

## Building the Client and Server Programs

```bash
git clone https://github.com/jaythemishra/golang-client-server
cd golang-client-server/dftMgr
go build
cd .. && cd dftNode
go build
```

## To build the linuxkit package and make a VM out of it

```bash
# Modify SERVER_URL in dftMgr/scripts/createClient.sh to be http://{YOUR_MACHINE'S_IP_ADDRESS}:12345/api/v1
cd dftNode
env GOOS=linux GOARCH=amd64 go build # This creates a linux executable
cd ..
linuxkit pkg push -org={YOUR_DOCKER_USERNAME} -hash=latest -disable-content-trust .
# Change dftNode's image in linuxkit.yml to be {YOUR_DOCKER_USERNAME}:latest-amd64
linuxkit build -format raw-bios linuxkit.yml
qemu-img convert -f raw -O qcow2 linuxkit-bios.img linuxkit-dft-node-test
```

## To-Do

- Create UI for assigning jobs
- Add more functionality to dftNode