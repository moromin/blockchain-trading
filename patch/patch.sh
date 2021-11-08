#!/bin/sh

# Configure gopatch
GOPATCH="gopatch"
OPTIONS="-p"

json_file=`cat $1`
json_length=`echo $json_file | jq length`

# Set file pass
CW_PKG_PATH="$GOPATH/pkg/mod/code.cryptowat.ch/cw-sdk-go@v1.0.3-0.20200521131807-bba79f5fb34f"
json_file=`echo $json_file | sed -e "s|CW_PKG_PATH|$CW_PKG_PATH|g"`

#Run gopatch
for i in `seq 0 $(expr $json_length - 1)`
do
	source_file=`echo $json_file | jq -r .[$i].source`
	patch_array=`echo $json_file | jq -r .[$i].patch`
	patch_length=`echo $json_file | jq ".[$i].patch | length"`
	chmod 644 $source_file

	for j in `seq 0 $(expr $patch_length - 1)`
	do
		patch_file=`echo $patch_array | jq -r .[$j]`
		echo $GOPATCH $OPTIONS $patch_file $source_file
		$GOPATCH $OPTIONS $patch_file $source_file
	done
	
done
