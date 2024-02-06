#!/bin/sh

export PATH=$PATH:/usr/local/go/bin:/home/$(logname)/go/bin

# source path for current session
. $HOME/.profile

GO_VERSION=1.21.6


inf() {
    echo "${2}\033[1;32m${1}\033[0m${2}"
}

wrn() {
    echo "${2}\033[1;33m${1}\033[0m${2}"
}

err() {
    echo "\n\033[1;31m${1}\033[0m\n"
    if [ ! -z ${2} ]
    then
        exit ${2}
    fi
}


get_go() {
    which go > /dev/null 2>&1 && return
    inf "Installing Go ${GO_VERSION} ..."
    # install golang per https://golang.org/doc/install#tarball
    curl -kL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | sudo tar -C /usr/local/ -xzf - \
    && echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> $HOME/.profile \
    && . $HOME/.profile \
    && go version
}


run_fp() {
    get_go
    oldpwd=${PWD}
    if [ $# -ne 1 ]; then
        inf "run fp"
        err "Usage: run_fp [testpath]"
        return 1
    fi
    cd ${oldpwd} \
    && cd testbed \
    && inf "testbed reservation in process" \
    && CGO_ENABLED=0 go run -v testbed.go \
    && inf "testbed reservation done"

    inf "running featureprofile test"
    cd ${FEATUREPROFILES_HOME} \
    && cp ${oldpwd}/testbed/otg.binding ${FEATUREPROFILES_HOME}/topologies/. \
    && CGO_ENABLED=0 go test -v "${1}" -binding "${FEATUREPROFILES_HOME}/topologies/otg.binding" 
}

help() {
    inf "Welcome to testbedops - The easisest way to manage opentestbed operations"
    wrn "Usage ./testbedops.sh [subcommand]: "
    echo "run_fp [testpath]  - Execute given test from featureprofiles"
    echo "\n"
}

case $1 in
    ""  )
        err "usage: $0 [name of any function in script]" 1
    ;;
    *   )
        # shift positional arguments so that arg 2 becomes arg 1, etc.
        cmd=${1}
        shift 1
        ${cmd} ${@} || err "failed executing './testbedops.sh ${cmd} ${@}'"
    ;;
esac
