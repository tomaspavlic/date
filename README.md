# date.Date

[![Tests](https://github.com/tomaspavlic/date/actions/workflows/go.yml/badge.svg)](https://github.com/tomaspavlic/date/actions/workflows/go.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tomaspavlic/date)
[![Go Reference](https://pkg.go.dev/badge/github.com/tomaspavlic/date.svg)](https://pkg.go.dev/github.com/tomaspavlic/date)
![GitHub](https://img.shields.io/github/license/tomaspavlic/date)

Represents dates with values from January 1, 0001 Anno Domini (Common Era) through December 31, 9999 A.D (C.E) in Gregorian calendar. `date.Date` represents number of days since January 1, 001.

## Installation
`date.Date` is compatible with modern Go releases in module mode, with Go installed:

```
go get github.com/tomaspavlic/date
```

## date.Date vs time.Time

`time.Time` consists of wall and ext encode the wall time seconds, wall time nanoseconds, and optional monotonic clock reading in nanoseconds. `time.Time` also contains `time.Location`. The overall size of the struct is 24 bytes. Working in `date.Date` is reduced to just be just single int64 therefore 1/3 of `time.Time` size.

## Usage

```golang
import "github.com/tomaspavlic/date"
```

```golang
// Create date.Date by just year, month and day
someDate := date.Create(2000, 1, 1)
// `date.Date` has similar functions as build-in `time.Time`
numberOfDays := date.Since(someDate)
fmt.Println("Number of days since 2000-01-01:", numberOfDays)
```

## More
Documentation can be found [on pkg.go.dev](https://pkg.go.dev/github.com/tomaspavlic/date).

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
