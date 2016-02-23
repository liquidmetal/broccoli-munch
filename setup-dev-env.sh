#!/bin/bash

export GOPATH=`pwd`
export PATH=`pwd`/bin:/usr/local/go/bin/:$PATH
export GOROOT=/usr/local/go/


###############
# WARNING
###############
# Do not commit any modifications to git when you add an api key here

# A bunch of API keys. You would have to set these up yourself.
export APIKEY_YOUTUBE=bleh
export APIKEY_GITHUB=bleh
export APIKEY_TWITTER=bleh
