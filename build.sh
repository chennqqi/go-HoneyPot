#!/usr/bin/env bash
path=$(dirname $0)
dir=$(cd $path && pwd)
export GOPATH=$dir:$GOPATH
app=hpot

mkdir -p bin

function compress() {
	upxexist=`ls tools|grep upx|wc -l`
	if [ $upxexist == "0" ] ; then
		wget -O upx.tar.xz https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz
		tar xvf upx.tar.xz -C tools
		mv tools/upx-3.95-amd64_linux/upx tools/
	fi
	./tools/upx -9 "bin/$app"
}

param=$1
case $param in
    debug)
	go build -o "bin/$app"
        ;;
    release)
	go build -o "bin/$app" -ldflags "-w -s"
	compress
        ;;
    *)
	echo "invalid argument, plz input [debug/release]"
esac

