#!/bin/bash

LATEST_WORKS=https://openlibrary.org/data/ol_dump_works_latest.txt.gz
LATEST_AUTHORS=https://openlibrary.org/data/ol_dump_authors_latest.txt.gz

wget -o ol_dumps_download_log -O latest_works.txt.gz $LATEST_WORKS
wget -a ol_dumps_download_log -O latest_authors.txt.gz $LATEST_AUTHORS

gzip -d latest_works.txt.gz
gzip -d latest_authors.txt.gz

