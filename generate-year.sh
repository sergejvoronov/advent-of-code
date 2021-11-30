#!/bin/bash

set -e

YEAR=${1}
CURRENT_DIR=$PWD
ROOT_DIR="advent-of-code"
INTERNAL_YEAR_CODE_DIR="internal/year"
CURRENT_DIR=$(echo $PWD | rev | cut -d '/' -f 1 | rev)

ALREADY_EXISTS="already exists. Skipping."
SUCCESSFULLY_CREATED="successfully created."

if [ -z "${YEAR}" ]; then
    echo "Error: year is not provided."
    exit 1
fi

if [[ "${CURRENT_DIR}" != "${ROOT_DIR}" ]]; then
    echo "Error: script generator must be executed from \"${ROOT_DIR}\" (project root) directory."
    exit 1
fi

generate_year_directory() {
    if [[ ! -d "${INTERNAL_YEAR_CODE_DIR}/${YEAR}" ]]; then
        mkdir -p "${INTERNAL_YEAR_CODE_DIR}/${YEAR}";
        echo "${CURRENT_DIR}/${INTERNAL_YEAR_CODE_DIR}/${YEAR} directory ${SUCCESSFULLY_CREATED}"
    else
        echo "${CURRENT_DIR}/${INTERNAL_YEAR_CODE_DIR}/${YEAR} directory ${ALREADY_EXISTS}"
    fi
}

generate_day_files() {
    for x in {1..25}
    do
        [[ ${#x} -gt 1 ]] && DAY_STRING=${x} || DAY_STRING="0${x}"

        if [[ ! -f "${INTERNAL_YEAR_CODE_DIR}/${YEAR}/day${DAY_STRING}.go" ]]; then
echo "package aoc${YEAR}

import (
	\"github.com/sergejvoronov/advent-of-code/internal/solution\"
)

type day${DAY_STRING} struct{}

func Day${DAY_STRING}() solution.Provider {
	return &day${DAY_STRING}{}
}

func (*day${DAY_STRING}) SolveA(input string) string {
	return solution.NotImplemented
}

func (*day${DAY_STRING}) SolveB(input string) string {
	return solution.NotImplemented
}" > "${INTERNAL_YEAR_CODE_DIR}/${YEAR}/day${DAY_STRING}.go"
            echo "${CURRENT_DIR}/${INTERNAL_YEAR_CODE_DIR}/${YEAR}/day${DAY_STRING}.go file ${SUCCESSFULLY_CREATED}"
        else
            echo "${CURRENT_DIR}/${INTERNAL_YEAR_CODE_DIR}/${YEAR}/day${DAY_STRING}.go file ${ALREADY_EXISTS}"
        fi
    done
}

generate_year_file() {
if [[ ! -f "${INTERNAL_YEAR_CODE_DIR}/${YEAR}/year.go" ]]; then
echo "package aoc${YEAR}

import \"github.com/sergejvoronov/advent-of-code/internal/solution\"

var Solutions = []solution.Provider{
	Day01(),
	Day02(),
	Day03(),
	Day04(),
	Day05(),
	Day06(),
	Day07(),
	Day08(),
	Day09(),
	Day10(),
	Day11(),
	Day12(),
	Day13(),
	Day14(),
	Day15(),
	Day16(),
	Day17(),
	Day18(),
	Day19(),
	Day20(),
	Day21(),
	Day22(),
	Day23(),
	Day24(),
	Day25(),
}" > "${INTERNAL_YEAR_CODE_DIR}/${YEAR}/year.go"
    echo "${CURRENT_DIR}/${INTERNAL_YEAR_CODE_DIR}/${YEAR}/year.go file ${SUCCESSFULLY_CREATED}"
else
    echo "${CURRENT_DIR}/${INTERNAL_YEAR_CODE_DIR}/${YEAR}/year.go file ${ALREADY_EXISTS}"
fi
}

generate_year_directory
generate_day_files
generate_year_file
