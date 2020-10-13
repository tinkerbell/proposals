#!/usr/bin/env bash

set -eu

d=$(cd "$(dirname "$0")/../proposals" && pwd)
proposal=$d/new/README.md

echo "🐙 >>> Creating a new proposal"

if [[ -f "$proposal" ]]; then
	echo "😞 >>> A proposal already exists in $f. Do you need to submit that one first?"
	exit 1
fi

mkdir -p "$d/new"
cp "$(dirname "$0")/README.md" "$proposal"

echo "🍀 >>> Your proposal is all set and ready in $proposal. Good luck!"
