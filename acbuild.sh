#!/bin/bash

set -ex

ACI_FILENAME=hello-latest-linux-$(arch).aci

rm -f $ACI_FILENAME

acbuild begin
acbuild set-name example.com/hello
acbuild copy hello /bin/hello
acbuild set-exec /bin/hello
acbuild write $ACI_FILENAME
acbuild end
