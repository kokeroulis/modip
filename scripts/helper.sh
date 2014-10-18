#!/usr/bin/env bash

gopath=$PWD/.gopath
gobin=$PWD/bin
p=$gopath/src/github.com/kokeroulis

update_deps() {
    github_deps=("codegangsta/negroni"
                 "mholt/binding"
                 "gorilla/mux"
                 "gorilla/sessions"
                 "boj/redistore"
                 "lib/pq"
                 "jacobsa/oglematchers"
                 "vaughan0/go-ini"
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

install_go() {
    if [ -d $gopath ]; then
        echo "Please remove $gopath"
       # exit 1
    fi

    mkdir -p $gopath/src/github.com
    mkdir -p $gopath/src/gopkg.in

    ln -sfv $PWD/vendor/github.com/* $gopath/src/github.com
    ln -sfv $PWD/vendor/gopkg.in/* $gopath/src/gopkg.in

    p=$gopath/src/github.com/kokeroulis/modip

    if [ ! -d $p ]; then
        mkdir -p $p
    fi

    my_packages=("api"
                 "db"
                 "models"
                 "config"
    )

    for i in ${my_packages[@]}; do
        ln -sfv $PWD/$i $p/$i
    done
}

install_js() {
    if [ -d node_modules ]; then
        echo "Please remove node_modules"
        exit 1
    fi

    if [ -d public/bower_components/angular ]; then
        echo "Please remove public/bower_components/"
        exit 1
    fi

    npm install
    ./node_modules/bower/bin/bower install
}

install() {
    install_go
    install_js
}

clean() {
    arch=$(go version | awk '{print $4}')
    rm -rf $gopath/pkg/$arch/github.com/kokeroulis/modip
}

run() {
    clean
    GOPATH=$gopath go run modip.go --start
}

function readConfigurationValue()
{
    local key=$1
    awk -v key="$key" -F "=" '$0 ~ key  {print $2}' modip.conf
}

setup_db() {

    DATABASE_ADMIN=$(readConfigurationValue "database_admin")
    DATABASE_USER=$(readConfigurationValue "database_user")
    DATABASE=$(readConfigurationValue "database_name")

    if [ $1 == "development" ]; then
        dropdb -U $DATABASE_ADMIN $DATABASE
    else
        dropdb -i -U $DATABASE_ADMIN $DATABASE
    fi

    createdb -U $DATABASE_ADMIN -O $DATABASE_USER $DATABASE
    PGOPTIONS='--client-min-messages=warning' psql -U $DATABASE_ADMIN $DATABASE -f $PWD/db/schema/superUserCommands.sql

    clean
    GOPATH=$gopath go run modip.go --setup-db=true
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
elif [ $1 == "install-go" ]; then
    install_go
elif [ $1 == "install-js" ]; then
    install_js
elif [ $1 == "run" ]; then
    run
elif [ $1 == "testdata" ]; then
    testdata
elif [ $1 == "db-create" ]; then
    setup_db $2
fi

