[![GoDoc](https://godoc.org/github.com/henderjon/durations?status.svg)](https://godoc.org/github.com/henderjon/durations)
[![License: BSD-3](https://img.shields.io/badge/license-BSD--3-blue.svg)](https://img.shields.io/badge/license-BSD--3-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/henderjon/durations)](https://goreportcard.com/report/github.com/henderjon/durations)
[![Build Status](https://travis-ci.org/henderjon/durations.svg?branch=dev)](https://travis-ci.org/henderjon/durations)
![tag](https://img.shields.io/github/tag/henderjon/durations.svg)
![release](https://img.shields.io/github/release/henderjon/durations.svg)

# durations

some helpful time/duration funcs

This package is mostly sugar. If you want to do duration math with properly typed ints (think return values from other funcs) you have to cast that int to a time.Duration (int64). This package makes the invocations a
bit more fluent and hides the casts from view.
