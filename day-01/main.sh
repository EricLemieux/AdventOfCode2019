#!/usr/bin/env bash

function find_fuel() {
  weight=${1}
  div=$((weight/3))
  final=$((div-2))
  echo "${final}"
}

function main() {
  sum=0

  # Read the input file and sum the result of all the elements
  while read -r line
  do
    fuel=$(find_fuel line)
    sum=$((sum+fuel))
  done < "${1:-/dev/stdin}"

  echo "${sum}"
}

main "${@}"
