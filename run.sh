#!/bin/bash

day=$(date +'%-d')
year=$(date +'%-Y')

while getopts y:d: flag
do
  case "${flag}" in
    y) year=${OPTARG};;
    d) day=${OPTARG};;
  esac
done


dayf=$(printf "%02d" $day)

go run calendar/$year/day-$dayf/day$dayf.go
