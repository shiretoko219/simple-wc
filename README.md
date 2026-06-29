# simple-wc

A beginner-friendly reimplementation of the Unix `wc` utility in Go.

## Overview

This project simplifies FreeBSD's `wc.c` into a single Go file for learning purposes. It counts lines, words, and bytes, matching the default behavior of the standard `wc` tool.

## Build & Run

```bash
go build -o my_wc main.go
./my_wc test.txt
```

## Acknowledgment

This project borrows the some logic from FreeBSD's `wc.c`. The original license is included below.
This is a learning exercise. It is not useful for real work — just use the real wc.
This project is licensed under BSD 3-Clause License, same as the referenced FreeBSD `wc.c`
