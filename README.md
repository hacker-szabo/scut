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

TODO

# Building

TODO

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