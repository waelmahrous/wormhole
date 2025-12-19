#!/usr/bin/env bash

destination=a_folder
file=a_file
wormhole_id=demo

rm -rf $destination $file
mkdir $destination
echo 'hello world' > $file

wormhole --id $wormhole_id open --destination $destination
vhs ./send.tape

rm -rf $destination
