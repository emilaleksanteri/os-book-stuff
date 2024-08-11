#!/usr/bin/env bash

my_var="100"

change_var () {
	echo "${1}"
	echo "${my_var}"
	my_var="$1"
	echo "${my_var}"
}

echo "${my_var}"
change_var 10 & change_var 1000

echo "hello" & echo "goodbye"
