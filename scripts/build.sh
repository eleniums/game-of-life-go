#!/bin/bash
set -e

NAME=gameoflife
TARGETS=windows/amd64,darwin/amd64 # https://github.com/karalabe/xgo#limit-build-targets
INSTALL_DEPENDENCIES=${INSTALL_DEPENDENCIES:-false}
VERSION=$(git describe --tags --always --long --dirty)
BASE=$(pwd)
RELEASE=$BASE/release/$VERSION

echo "Creating release of $NAME with version: $VERSION"

# build all binaries
echo "Using xgo for cross platform build: https://github.com/karalabe/xgo"
if $INSTALL_DEPENDENCIES = true
then
    echo "Retrieving latest xgo..."
    docker pull karalabe/xgo-latest
    go get -u github.com/karalabe/xgo
fi
xgo --targets=$TARGETS --out $NAME github.com/eleniums/game-of-life-go/cmd/game

# create release folder
rm -rf $RELEASE
mkdir -p $RELEASE

# package windows binary
if [[ $TARGETS = *"windows/amd64"* ]]
then
    echo "Packaging Windows binary..."
    mv $(ls | grep $NAME-windows) $RELEASE/$NAME.exe
    cd $RELEASE
    zip "$NAME"_windows.zip $NAME.exe
    rm -rf $NAME.exe
    cd $BASE
fi

# package linux binary
if [[ $TARGETS = *"linux/amd64"* ]]
then
    echo "Packaging Linux binary..."
    mv $(ls | grep $NAME-linux) $RELEASE/$NAME
    cd $RELEASE
    tar -czvf "$NAME"_linux.tar.gz $NAME
    rm -rf $NAME
    cd $BASE
fi

# package mac binary
if [[ $TARGETS = *"darwin/amd64"* ]]
then
    echo "Packaging macOS binary..."
    mv $(ls | grep $NAME-darwin) $RELEASE/$NAME
    cd $RELEASE
    tar -czvf "$NAME"_mac.tar.gz $NAME
    rm -rf $NAME
    cd $BASE
fi

echo "Release of $NAME created: $RELEASE"