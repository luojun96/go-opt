#!/bin/bash

path=$1

if [ -z "$path" ]; then
    echo "Usage: $0 path"
    exit 1
fi

if [ ! -d "$path" ]; then
    echo "Path $path does not exist"
    exit 1
fi

echo "Changing package name in $path"
directories=$(find $path -type d)

# Loop through all directories and change package name
for directory in $directories; do
    target=$(basename $directory)
    echo "Changing package name in $directory to $target"
    pushd $directory
    # cmd: find ./ -type f -exec sed -i '' -e "s/main/$(basename $(pwd))/g" {} \;
    # https://stackoverflow.com/questions/19456518/error-when-using-sed-with-find-command-on-os-x-invalid-command-code
    find $directory -type f -name '*.go' -exec sed -i '' -e "s/main/$target/g" {} \;
    popd
done

