#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/demothrift"
exec "$CURDIR/bin/demothrift"