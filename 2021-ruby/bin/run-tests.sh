#!/bin/bash

o=$(pwd)

for d in $(find . -name spec -type d | sed 's/\/spec/ /'); do
  cd "$d"
  echo $(pwd)
#   if [[ "$d" == "day-23" ]]; then
#     echo "Skipping day-23 as it is very slow"
#     continue
#   fi
  echo "BUNDLING..."
  bundle update --bundler
  bundle install >/dev/null 2>&1
  echo "TESTING..."
  bundle exec rspec -f d
  echo "DONE."
  cd "$o"
done

cd $o > /dev/null
