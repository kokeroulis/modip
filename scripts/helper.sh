#!/usr/bin/env bash

gopath=$HOME/.modip-build
gobin=$PWD/bin
p=$gopath/src/github.com/kokeroulis


install() {
    if [ -d $gopath ]; then
        echo "Please remove $gopath"
        exit 1
    fi

    github_deps=("codegangsta/negroni"
                 "goincremental/negroni-sessions"
                 "mholt/binding"
                 "gorilla/mux"
                 "lib/pq"
                 "smartystreets/goconvey"
                 "jacobsa/oglematchers"
    )

    gopkg_deps=("unrolled/render.v1")

    for i in ${github_deps[@]}; do
        echo "Installing: $i"
        GOPATH=$gopath go get github.com/$i
    done

    for i in ${gopkg_deps[@]}; do
        echo "Installing: $i"
        GOPATH=$gopath go get gopkg.in/$i
    done

    p=$gopath/src/github.com/kokeroulis

    if [ ! -d $p ]; then
        mkdir -p $p
    fi

    ln -sf $PWD $p/modip
}

run() {
    GOPATH=$gopath go run modip.go
}

testdata() {
    GOPATH=$gopath go run tools/modip-importer.go --data=testdata.json
}

if [ -z $1 ]; then
    echo "Invalid arguments"
    exit 1
fi

if [ $1 == "install" ]; then
    install
elif [ $1 == "run" ]; then
    run
elif [ $1 == "test" ]; then
    pushd $p > /dev/null
    GOPATH=$gopath go test -v  -coverprofile=c.out github.com/kokeroulis/modip/tests
    popd > /dev/null
elif [ $1 == "testdata" ]; then
    testdata
elif [ $1 == "test-web" ]; then
    GOPATH=$gopath $gopath/bin/goconvey
elif [ $1 == "cover" ]; then
    GOPATH=$gopath go tool cover -html=$p/c.out -o .coverage.html
    xdg-open $PWD/.coverage.html
fi

