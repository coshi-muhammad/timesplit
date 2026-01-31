#!/usr/bin/zsh

pushd ./cmd/timesplit/ 
cp ../../Icon.png ./
fyne package -os linux 
cp timesplit.tar.xz ../../bin/linux
rm Icon.png
popd


