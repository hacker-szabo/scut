# Superb Cut

UNDER DEVELOPMENT!!!

A clone for the usual linux "cut" command but works with multiple characters provided to the -d parameter.

SCut means Superb Cut because it's superb.

# Installation

TODO

# Building

TODO

# Flags

```
scut [-b list] [-c list] [-f list] [-n] [-d delim] [-s] [file]
```

## -b

Bytes; a list following -b specifies a range of bytes which will be returned, e.g. cut -b1-66 would return the first 66 bytes of a line. NB If used in conjunction with -n, no multi-byte characters will be split. NNB. -b will only work on input lines of less than 1023 bytes

## -c

Characters; a list following -c specifies a range of characters which will be returned, e.g. cut -c1-66 would return the first 66 characters of a line

## -f

Specifies a field list, separated by a delimiter
list
A comma separated or blank separated list of integer denoted fields, incrementally ordered. The - indicator may be supplied as shorthand to allow inclusion of ranges of fields e.g. 4-6 for ranges 4–6 or 5- as shorthand for field 5 to the end, etc.

## -n

Used in combination with -b suppresses splits of multi-byte characters

## -d

Delimiter; the character immediately following the -d option is the field delimiter for use in conjunction with the -f option; the default delimiter is tab. Space and other characters with special meanings within the context of the shell in use must be enquoted or escaped as necessary.

## -s

Bypasses lines which contain no field delimiters when -f is specified, unless otherwise indicated.


## file

The file (and accompanying path if necessary) to process as input. If no file is specified then standard input will be used.