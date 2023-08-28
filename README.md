# Superb Cut

UNDER DEVELOPMENT!!!

Implements a missing feature of the "cut" common tools linux/unix command which is: **-d accepts a string not just a single character**

SCut means Superb Cut because it's superb.

# Installation

TODO

# Building

TODO

# Flags

```
scut [-d delimeter_string] [-f fragment 2] [-f fragment 2] [file]
```

## -d

Delimiter string; the string provided with the -d option is the delimeter where the rows will be cut; the default delimiter is a tab character ("\t"). Space and other characters with special meanings within the context of the shell in use must be enquoted or escaped as necessary.

## -f

The fragment number (starts with 1, not 0) to print each line.
The same as -f in the original cut command but you can add multiple of them and they will be
printed all with a TAB character in between.

## file

The file (and accompanying path if necessary) to process as input. If no file is specified then standard input will be used.

## Usage

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