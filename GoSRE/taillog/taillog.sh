#!/bin/bash

numberRegex='^[0-9]+([.][0-9]+)?$'

#today=$(date '+%Y-%m-%d,%H:%M:%S.%N')
#currentMinutes=$(date +%s)/60

# minutes in a year 525600
# minutes in a month 43800
# minutes in a day 1440

timeRange=""
iVar="8.8.8.8"
dVar=""

fileName=${@: -1}


seconds=$(awk -F: '{ print ($1 * 60) + $2 + ($3 / 60) }' <<< "00:20:40")
echo $seconds


if [[ -z $fileName ]] ; then
	echo "Need to pass a csv file"
	exit 1
fi

while getopts ":t:id" opt 
do
	case ${opt} in
		t) 
			if ! [[ ${OPTARG} =~ $numberRegex ]] ; then
				echo "error: Not a number" >&2; exit 1
			else
				timeRange=${OPTARG}
			fi
			;;

		i) iVar=${OPTARG};;
		d) dVar=${OPTARG};;
		\?) echo "Invalid option: -$OPTARG" >&2
			exit 1;;
		:) echo "Option -$OPTARG requires an argument." >&2
		   exit 1;;
	esac
done

if [[ -z "$timeRange" ]] ; then
	echo "Need to set option: -t" >&2
	exit 1
fi

# Set to the number of the last line first
lineNumber=$(wc -l < ${fileName})

#sed -n -e '$p' $fileName
line=$(sed "${numberOfLines}q;d" $fileName)

