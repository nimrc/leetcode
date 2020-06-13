#!/bin/bash

function create_new_question()
{
	if [[ -n "$1" && -n "$2" ]]
	then
		echo $1
		q=`echo $2 | cut -d '/' -f5`
		src=${q//'-'/'_'}
		mkdir -p "src/$1_$src"
		echo "package solution" > "src/$1_$src/$src.go"
		echo "package solution" > "src/$1_$src/${src}_test.go"
	else
		echo "Error: missing required parameters."
		echo "Usage: "
		echo "  ./create.sh question_no url"
		echo "Example:"
		echo "  ./create.sh 001 https://leetcode.com/problems/two-sum/"
	fi
}

create_new_question $1 $2
