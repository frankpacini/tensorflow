# TensorFlow in Go
Construct and execute TensorFlow graphs in Go.

[![GoDoc](https://godoc.org/github.com/frankpacini/tensorflow/tensorflow/go?status.svg)](https://godoc.org/github.com/frankpacini/tensorflow/tensorflow/go)

## Quickstart:
1. Install the Tensorflow C library
```bash
curl -L https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.5.0.tar.gz | sudo tar xz --directory /usr/local
sudo ldconfig
```
2. Add the module
```bash
# cd into telegraf-mirror/
go mod edit -require github.com/frankpacini/tensorflow@v0.0.0-20210721220828-30b5537d0625
```
Here the version number is based on the date and SHA of the latest commit. Proper versioning can be set up later. For now if new commits are added, the version number can be determined by getting the timestamp of the commit in UTC (hover over the relative commit time) and the first 12 characters of the SHA.


## Creating or updating this package:
1. Install the Tensorflow C library
```bash
curl -L https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.5.0.tar.gz | sudo tar xz --directory /usr/local
sudo ldconfig
```
2. Install protobuf dependencies
```bash
sudo apt install libprotobuf-dev
```
```bash
# Protobuf install (may take a while):
cd ~/Downloads
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-all-3.6.1.tar.gz
tar -xvzf protobuf-all-3.6.1.tar.gz 
cd protobuf-3.6.1/
./configure
make
make check
sudo make install
sudo ldconfig
```
3. Clone Tensorflow source with fixes for Go
```bash
# Change the branch as necessary. Clone wherever convenient
sudo git clone --branch r2.5-go https://github.com/galeone/tensorflow.git ~/go/src/github.com/galeone/tensorflow
cd ~/go/src/github.com/galeone/tensorflow
```
4. Add Go to Sudo
```bash
# First copy go binary folder location. Find this with "which go"
sudo visudo
# Copy folder location into the secure_path variable with a colon separating it from the previous entry
```
5. Setup go modules
```bash
sudo go mod init github.com/galeone/tensorflow
cd tensorflow/go
sudo go mod vendor
```
6. Generate proto files and TF ops wrappers
```bash
sudo go generate ./...
sudo go mod tidy
```
7. Change username references
```bash
cd ../..
find ./ -type f -exec sed -i -e 's/galeone/NEW_USERNAME/g' {} \;
```
8. Push new repository
```bash
sudo git remote set-url NEW_URL
sudo git push
```
