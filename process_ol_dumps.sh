#!/bin/bash

source /home/elvis/.rails_envs # make sure this is set in production

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

# source then call it from dir where recent files are for manual import
import_resources() {
  RECENT_WORKS="$(pwd)/recent_works.txt"
  RECENT_AUTHORS="$(pwd)/recent_authors.txt"

  cd $WYSEBITS_API_DIR
  bundle exec rake db:import_books[$RECENT_WORKS]
  bundle exec rake db:import_authors[$RECENT_AUTHORS]
}

cleanup_dumps() {
  rm latest_*.txt
}

cleanup_recent_files() {
  rm recent_*.txt
}

# run script
download_dumps
wait
extract_dumps
wait

if [ -z "$1" ]; then
  extract_recent_resources
  wait
else
  extract_recent_resources_with_timestamp $1
  wait
fi

cleanup_dumps
wait
#cleanup_recent_files;
