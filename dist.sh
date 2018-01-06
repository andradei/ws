#!/bin/sh
os=( 'darwin' 'linux' )
arch=('amd64' '386' )

for goos in ${os[@]}
do
  for goarch in ${arch[@]}
  do
    GOOS=$goos GOARCH=$goarch go build -o "bin/ws_${goos}_${goarch}"
  done
done
