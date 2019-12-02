#!/usr/bin/env bash

function find_fuel() {
  local weight=${1}
  local div=$((weight/3))
  local fuel=$((div-2))
  local fuel_for_fuel=0

  if [[ $fuel -gt 0 ]];
  then
    fuel_for_fuel=$(find_fuel $fuel)
  fi

  local final=$((fuel+fuel_for_fuel))

  if [[ $final -lt 0 ]];
  then
    final=0
  fi

  echo "${final}"
}

function main() {
  local sum=0

  # Read the input file and sum the result of all the elements
  while read -r line
  do
    local fuel=0
    fuel=$(find_fuel line)
    sum=$((sum+fuel))
  done < "${1:-/dev/stdin}"

  echo "${sum}"
}

main "${@}"
