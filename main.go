/*
 * SPDX-License-Identifier: BSD 3-Clause License
 * Copyright (c) 2026 K.Shiretoko
 *
 * Portions of this file are derived from FreeBSD wc.c:
 *   Copyright (c) 1980, 1987, 1991, 1993
 *       The Regents of the University of California.
 *   SPDX-License-Identifier: BSD-3-Clause
 *   Full text: see LICENSE.freebsd
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	tlinect int64
	twordct int64
	tcharct int64

	doline bool
	doword bool
	dochar bool
)

func main() {
	flag.BoolVar(&doline, "l", false, "count lines")
	flag.BoolVar(&doword, "w", false, "count words")
	flag.BoolVar(&dochar, "c", false, "count characters")
	flag.Parse()

	if !(doline || doword || dochar) {
		doline = true
		doword = true
		dochar = true
	}

	files := flag.Args()

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Must input a file.")
		os.Exit(1)
	}

	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", filename, err)
			os.Exit(1)
		}
		linect, wordct, charct := cnt(f)
		showct(linect, wordct, charct, filename)

		tlinect += linect
		twordct += wordct
		tcharct += charct

		f.Close()
	}
	if len(files) > 1 {
		showct(tlinect, twordct, tcharct, "total")
	}
}

func showct(linect, wordct, charct int64, file string) {
	if doline {
		fmt.Printf(" %d", linect)
	}
	if doword {
		fmt.Printf(" %d", wordct)
	}
	if dochar {
		fmt.Printf(" %d", charct)
	}
	fmt.Printf(" %s\n", file)
}

func cnt(f *os.File) (int64, int64, int64) {
	var linect, wordct, charct int64
	inWord := false

	reader := bufio.NewReader(f)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, 0
		}
		charct++
		if b == '\n' {
			linect++
		}
		isSpace := b == ' ' || b == '\t' || b == '\n' || b == '\r' || b == '\f' || b == '\v'

		if isSpace {
			inWord = false
		} else if !inWord {
			wordct++
			inWord = true
		}
	}
	return linect, wordct, charct
}
