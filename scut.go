package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IntSlice []int

func (i *IntSlice) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *IntSlice) Set(value string) error {
	val, err := strconv.Atoi(value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid value for -f: %s\n", value)
		os.Exit(1)
	}
	*i = append(*i, val)
	return nil
}

type Arguments struct {
	delimiter       string
	file            string
	fragments       IntSlice
	columnSeparator string
	// skipInvalidRows bool
}

func readFlags() Arguments {
	ret := Arguments{}
	flag.StringVar(&ret.delimiter, "d", "\t", "delimiter used for splitting each line")
	flag.Var(&ret.fragments, "f", "Fragments to extract from each line")
	// flag.BoolVar(&ret.skipInvalidRows, "s", false, "Skip invalid rows")
	flag.StringVar(&ret.columnSeparator, "cs", "\t", "delimiter used for splitting each column")

	flag.Parse()
	// reading positional arguments (file)
	positionalArguments := flag.Args()
	if flag.NArg() > 0 {
		ret.file = positionalArguments[0]
	}

	return ret
}

func printLine(line string, fragments IntSlice, delimiter string, columnSeparator string,

// TODO implement this
// skipInvalidRows bool

) {
	if len(line) == 0 {
		return
	}

	if len(fragments) == 0 {
		fragments = append(fragments, 1)
	}
	splitted := strings.Split(line, delimiter)
	// I am going to leave this here as a protest against
	// the tyrant go compiler and its extremist creators that
	// "promote" clean and readable code so much that they
	// go as far as to slow down debugging of good tax paying citizens
	// that just want to get their job done or learn a new language
	// which has a great reputation despite people have been
	// complaining about this "FEATURE" for over 10 years now or since the
	// language was created. Yes, I am talking about the ""
	_ = splitted

	for i := 0; i < len(fragments); i++ {
		f := fragments[i] - 1

		if f >= len(splitted) || f < 0 {
			fmt.Fprintf(os.Stderr, "\nInvalid fragment: %d\n(For line: %s)\n", f, line)
			os.Exit(1)
		}

		if i > 0 {
			fmt.Print(columnSeparator)
		}
		fmt.Print(splitted[f])
	}
	fmt.Println()
}

func readAndPrintInputData(args Arguments) {

	if args.file == "" {
		stdin, err := io.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}

		str := string(stdin)
		lines := strings.Split(str, "\n")
		for _, line := range lines {
			printLine(line, args.fragments, args.delimiter, args.columnSeparator)
		}
	} else {
		// read lines from file
		f, err := os.Open(args.file)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		i := 0
		for scanner.Scan() {
			printLine(scanner.Text(), args.fragments, args.delimiter, args.columnSeparator)
			i++
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}

func main() {
	args := readFlags()
	readAndPrintInputData(args)
	// fmt.Println(args)
}
