# Superb Cut

Implements a missing feature of the "cut" common tools linux/unix command which is: **-d accepts a string not just a single character**.

The tool also accepts multiple -f options.

Just as simple as that.

SCut means Superb Cut because it's superb.

Version: v1.0

## Feature I am considering

1. Program will fail if you ask for a fragment (-f) too high and finds a row that does not have that many fragments. I am considering a --skip-invalid-rows or something.

## Motivation

1. I wanted to learn GO, this is my second program implemented right after Hello World so please forgive me for stupid mistakes but please free to give me suggestions :)
2. I could not believe that cut still did not accept multiple characters for the -d parameter.

## Usage

Keep in mind: if you provide a filename, it must be the last parameter!

### Read a file

```
scut -d ";;;" -f 1 FILENAME
```

### Pipe to STDIN

```
cat long_file.txt | scut -d ";;;" -f 1
```

### Print multiple fragments

```
scut -d ";;;" -f 1 -f 3 FILENAME
```

### Print multiple fragments and change the column separator (TAB by default)

```
scut -d ";;;" -f 1 -f 3 -cs "," FILENAME
```

In special cases this can be used to change a sequence of characters (-d) to another sequence of characters (-cs).

# Installation

## Releases

I cross compiled without testing to the following operating systems and architectures:

- Windows 32/amd64/arm64
- Linux 32/amd64/arm64
- Macos amd64/arm64

## Go install

```
go install github.com/hacker-szabo/scut@latest
```

# Building

I used go 1.21 but it might work for way lower versions as well.

1. Download the source code (or at least scut.go)
2. Navigate to the directory in the terminal/command line
3. `go build scut.go`

## Cross compile

With go installed you can cross compiled by following the basic GO cross compile way:

Find your target operating system (GOOS) and architecture (GOARCH) in the following table:

$GOOS | $GOARCH
----------------
| aix | ppc64 |
| android | 386 |
| android | amd64 |
| android | arm |
| android | arm64 |
| darwin | amd64 |
| darwin | arm64 |
| dragonfly | amd64 |
| freebsd | 386 |
| freebsd | amd64 |
| freebsd | arm |
| illumos | amd64 |
| ios | arm64 |
| js | wasm |
| linux | 386 |
| linux | amd64 |
| linux | arm |
| linux | arm64 |
| linux | loong64 |
| linux | mips |
| linux | mipsle |
| linux | mips64 |
| linux | mips64le |
| linux | ppc64 |
| linux | ppc64le |
| linux | riscv64 |
| linux | s390x |
| netbsd | 386 |
| netbsd | amd64 |
| netbsd | arm |
| openbsd | 386 |
| openbsd | amd64 |
| openbsd | arm |
| openbsd | arm64 |
| plan9 | 386 |
| plan9 | amd64 |
| plan9 | arm |
| solaris | amd64 |
| wasip1 | wasm |
| windows | 386 |
| windows | amd64 |
| windows | arm |
| windows | arm64 |

Table copied from: https://go.dev/doc/install/source#environment

# Flags

```
scut -d delimeter_string -f fragment1 [-f fragment2] [-cs colum_separator] [file]
```

## -d

Delimiter string; the string provided with the -d option is the delimeter where the rows will be cut; the default delimiter is a tab character ("\t"). Space and other characters with special meanings within the context of the shell in use must be enquoted or escaped as necessary.

## -f

The fragment number (starts with 1, not 0) to print each line.
The same as -f in the original cut command but you can add multiple of them and they will be
printed all with a TAB character in between.

## -cs
Column separator. By default TAB. This character (or string ;)) will appear in the output between multiple columns if and only you provided multiple fragments.

## file

The file (and accompanying path if necessary) to process as input. If no file is specified then standard input will be used.


# Benchmarks compared to cut

TODO