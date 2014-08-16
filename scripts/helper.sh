#!/usr/bin/env bash

gopath=$PWD/tmp

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
}

run() {
    GOPATH=$gopath:$PWD go run modip.go
}

build_tools() {
    GOPATH=$gopath GOBIN=$gopath/bin go install tools/modip-importer.go
}

testdata() {
    $gopath/bin/modip-importer --data=testdata.json
}

if [ -z $1 ]; then
    echo "Invalid arguments"
    exit 1
fi

if [ $1 == "install" ]; then
    install
elif [ $1 == "run" ]; then
    run
elif [ $1 == "build-tools" ]; then
    build_tools
elif [ $1 == "test" ]; then
    _gopath=$gopath:$PWD

    pushd src/ > /dev/null
    GOPATH=$_gopath go test -v tests
    popd > /dev/null
elif [ $1 == "testdata" ]; then
    testdata
fi

