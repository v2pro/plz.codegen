# wombat

[![Sourcegraph](https://sourcegraph.com/github.com/v2pro/wombat/-/badge.svg)](https://sourcegraph.com/github.com/v2pro/wombat?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/v2pro/wombat)
[![Build Status](https://travis-ci.org/v2pro/wombat.svg?branch=master)](https://travis-ci.org/v2pro/wombat)
[![codecov](https://codecov.io/gh/v2pro/wombat/branch/master/graph/badge.svg)](https://codecov.io/gh/v2pro/wombat)
[![rcard](https://goreportcard.com/badge/github.com/v2pro/wombat)](https://goreportcard.com/report/github.com/v2pro/wombat)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/v2pro/wombat/master/LICENSE)
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/v2pro/Lobby)

<img src="wombat.png" width="250">

binding &amp; validation &amp; functional goodies. 

implements api defined in `v2pro/plz/util` package

# Example

```golang
//go:generate go install github.com/v2pro/wombat/cmd/wombat-codegen
//go:generate $GOPATH/bin/wombat-codegen -pkg github.com/v2pro/wombat/example
func init() {
	generic.Declare(func() {
		plz.Max(int(0))
		plz.Max(float64(0))
		plz.Max(model.User{}, "Score")
	})
}

func Demo_max_min(t *testing.T) {
	should := require.New(t)
	should.Equal(3, plz.Max(1, 3, 2))
	should.Equal(float64(3), plz.Max(1.0, 3.0, 2.0))
	should.Equal(model.User{3}, plz.Max(
		model.User{1}, model.User{3}, model.User{2},
		"Score"))
}
```

replace `github.com/v2pro/wombat/example` with your package name. 

you need call `go generate` before compile