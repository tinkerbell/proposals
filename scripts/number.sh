#!/usr/bin/env bash

if (($# != 1)) || ! [[ $1 =~ ^[0-9]+$ ]]; then
	echo ='usage: number.sh pull_request_number' >&2
	exit 1
fi

d=$(cd "$(dirname "$0")/../proposals" && pwd)
proposal=$d/new/README.md

echo "🐙 >>> Alright, ready to number the new proposal"

if ! [[ -f "$proposal" ]]; then
	echo "😞 >>> Expected a proposal in $proposal, none found..." >&2
	exit 1
fi

n=$(printf "%04d" "$1")
if [[ -d "$d/$n" ]]; then
	echo "😞 >>> Hmm $proposal already exists..." >&2
	exit 1
fi

mkdir $d/$n
sed --in-place "s|id: ID|id: $n|" $proposal
git mv $proposal $d/$n/README.md

echo ">>> All done!"
