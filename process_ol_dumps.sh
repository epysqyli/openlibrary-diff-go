#!/bin/bash

LATEST_WORKS=https://openlibrary.org/data/ol_dump_works_latest.txt.gz
LATEST_AUTHORS=https://openlibrary.org/data/ol_dump_authors_latest.txt.gz

download_dumps() {
  wget -o ol_dumps_download_log -O latest_works.txt.gz $LATEST_WORKS
  wget -a ol_dumps_download_log -O latest_authors.txt.gz $LATEST_AUTHORS
}

extract_dumps() {
  gzip -d latest_works.txt.gz
  gzip -d latest_authors.txt.gz
}

extract_recent_resources() {
  ./main latest_works.txt recent_works.txt
  ./main latest_authors.txt recent_authors.txt
}

extract_recent_resources_with_timestamp() {
  ./main latest_works.txt recent_works.txt $1
  ./main latest_authors.txt recent_authors.txt $1
}

cleanup_dumps() {
  rm latest_*.txt
}

# run script
download_dumps;
extract_dumps;

if [ -z "$1" ]; then
  extract_recent_resources
else
  extract_recent_resources_with_timestamp $1
fi

cleanup_dumps
