package main

// linux cut command clone

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const version = "v0.1"

type RawArguments struct {
	c            string
	cProvided    bool
	d            string
	dProvided    bool
	f            []string
	fProvided    bool
	s            bool
	n            bool
	b            string
	bProvided    bool
	file         string
	fileProvided bool
}

// Custom argument parser because I want -c to work as a flag as well as -c<value>
func parseArguments() RawArguments {
	ret := RawArguments{}
	rawArgs := os.Args[1:]

	isPreviousNeeded := false
	previousFlag := ""

	// for i = 0 to len(rawArgs)
	for i := 0; i < len(rawArgs); i++ {
		if rawArgs[i][0] == '-' {
			// TODO: -b -10 doesnt work!
			// check if previous flag is needed

			if isPreviousNeeded {
				// second letter is numeric
				if rawArgs[i][1] >= '0' && rawArgs[i][1] <= '9' {
					// this means, that we ha something like: -24
					if previousFlag == "c" {
						ret.c = rawArgs[i]
						ret.cProvided = true
						isPreviousNeeded = false
						continue
					}

					if previousFlag == "b" {
						ret.b = rawArgs[i]
						ret.bProvided = true
						isPreviousNeeded = false
						continue
					}
				} else {
					// this means that -c does not have a value, set it to ""
					if previousFlag == "c" {
						ret.c = ""
						ret.cProvided = true
						isPreviousNeeded = false

						// check if this is the last argument so we set it to "" here
						if i == len(rawArgs)-1 {
							ret.c = ""
							ret.cProvided = true
							isPreviousNeeded = false
						}
						continue

					}
				}
			}

			// if isPreviousNeeded && previousFlag == "c" {
			// 	ret.c = ""
			// 	ret.cProvided = true
			// 	isPreviousNeeded = false
			// }

			if len(rawArgs[i]) == 2 {
				// setting up -n flag
				if rawArgs[i][1] == 'n' {
					ret.n = true
					isPreviousNeeded = false
					continue
				}

				// setting up -s flag
				if rawArgs[i][1] == 's' {
					ret.s = true
					isPreviousNeeded = false
					continue
				}

				// for every other parameter an input is needed
				previousFlag = rawArgs[i][1:]
				isPreviousNeeded = true
				continue
			} else {
				// it contains a value!
				if len(rawArgs[i]) > 2 {
					// it's -c<value>
					if rawArgs[i][1] == 'c' {
						ret.c = rawArgs[i][2:]
						ret.cProvided = true
						isPreviousNeeded = false
						continue
					} else if rawArgs[i][1] == 'd' {
						ret.d = rawArgs[i][2:]
						ret.dProvided = true
						isPreviousNeeded = false
						continue
					} else if rawArgs[i][1] == 'f' {
						ret.f = append(ret.f, rawArgs[i][2:])
						ret.fProvided = true
						isPreviousNeeded = false
						continue
					} else if rawArgs[i][1] == 'b' {
						ret.b = rawArgs[i][2:]
						ret.bProvided = true
						isPreviousNeeded = false
						continue
					} else {
						// this is a list like -b -4 => -b 0-4
					}
				}
			}
		} else {
			// let's see if something is needed
			if isPreviousNeeded {
				if previousFlag == "c" {
					ret.c = rawArgs[i]
					ret.cProvided = true
					isPreviousNeeded = false
					continue
				} else if previousFlag == "d" {
					ret.d = rawArgs[i]
					ret.dProvided = true
					isPreviousNeeded = false
					continue
				} else if previousFlag == "f" {
					ret.f = append(ret.f, rawArgs[i])
					ret.fProvided = true
					isPreviousNeeded = false
					continue
				} else if previousFlag == "b" {
					ret.b = rawArgs[i]
					ret.bProvided = true
					isPreviousNeeded = false
					continue
				}
			}

			// fmt.Println("File provided: ", ret.fileProvided)
			// it's not starting with -, so must be "file"
			if ret.fileProvided == true {
				// todo: error message on stderr
				// fmt.Println("Error: Only one file can be provided!")
				os.Exit(1)
			}

			if len(rawArgs[i]) > 0 {
				ret.file = rawArgs[i]
				ret.fileProvided = true
			}
		}
	}

	return ret
}

type Arguments struct {
	c            []int
	cProvided    bool
	d            string
	dProvided    bool
	f            []int
	fProvided    bool
	s            bool
	n            bool
	b            []int
	bProvided    bool
	file         string
	fileProvided bool
}

// -1 means end of line
// TODO, nem megy erre: go run scut.go  -c -f1
// TODO, nem irja be erre: go run scut.go  -c
func listParser(input string) []int {
	first := 0
	second := -1

	// if '-' in input
	dashInInput := false
	for i := 0; i < len(input); i++ {
		if input[i] == '-' {
			dashInInput = true
			break
		}
	}

	if dashInInput == false {
		fmt.Fprintf(os.Stderr, "Error: Invalid list input: %s\n", input)
		os.Exit(1)
	} else if len(input) == 0 {
		ret := []int{-1, -1}
		return ret
	}

	splitInput := strings.Split(input, "-")
	if len(splitInput[0]) > 0 {
		first, _ = strconv.Atoi(splitInput[0])
	}

	if len(splitInput[1]) > 0 {
		second, _ = strconv.Atoi(splitInput[1])
	}

	// TODO error handling?

	ret := []int{first, second}

	return ret
}

func getArgumentsFromRawArguments(rawArgs RawArguments) Arguments {
	ret := Arguments{}
	ret.fProvided = rawArgs.fProvided
	ret.f = make([]int, len(rawArgs.f))
	for i := 0; i < len(rawArgs.f); i++ {
		ret.f[i], _ = strconv.Atoi(rawArgs.f[i])
	}

	if rawArgs.cProvided {
		ret.cProvided = true
		ret.c = listParser(rawArgs.c)
	}

	if rawArgs.bProvided {
		ret.bProvided = true
		ret.b = listParser(rawArgs.b)
	}

	ret.n = rawArgs.n
	ret.s = rawArgs.s

	ret.dProvided = rawArgs.dProvided
	ret.d = rawArgs.d

	ret.file = rawArgs.file
	ret.fileProvided = rawArgs.fileProvided

	return ret
}

func main() {
	rawArguments := parseArguments()
	arguments := getArgumentsFromRawArguments(rawArguments)

	fmt.Println("Arguments: ", arguments)
}
