#!/bin/sh

# Sleep so indexing app will create index.txt and filenames.txt files used to query
sleep 5

exec "$@"
