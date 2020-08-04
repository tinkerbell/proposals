#!/bin/bash

set -eu

RFD_ID=$1
RFD_TITLE=$2

echo "🐙 >>> started the creation process for RFD ${RFD_ID} titled \"${RFD_TITLE}\""

branch_exists=$(git branch -rl *${RFD_ID})

if [[ -z ${branch_exists} ]]; then
    git checkout -b ${RFD_ID}
    mkdir $RFD_ID
    cp ../prototypes/prototype.md $RFD_ID/README.md
    echo "🍀 >>> You are all set! Good luck"
    exit 0
else
    echo "😞 >>> branch with name ${RFD_ID} already exists. Check out with your
    co-workers or pick a different id."
    exit 1
fi

exit 1
