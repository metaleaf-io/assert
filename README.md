[![Build Status](https://travis-ci.org/metaleaf-io/assert.svg)](https://travis-ci.org/metaleaf-io/assert)
[![GoDoc](https://godoc.org/github.com/metaleaf-io/assert/github?status.svg)](https://godoc.org/github.com/metaleaf-io/assert)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.12.0-green.svg)](https://golang.org/doc/devel/release.html#go1.12)

# assert

Lightweight assertion library based on the fluent interface from
[assertj](http://joel-costigliola.github.io/assertj/)

## Features

The matchers included in our `assert` library are fully compatible with, and
depend on the standard Go [testing package](https://golang.org/pkg/testing/).
These just add a little syntactic sugar on top of the familiar test patterns.

To use the example from the testing documentation, here is how one would
normally write a test in Go:

```cgo
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

With the matchers included in our `assert` package, one would write:

```cgo
import "github.com/metaleaf-io/assert"

func TestAbs(t *testing.T) {
    got := Abs(-1)
    assert.With(t).
        That(got).
        IsEqualTo(1)
}
```

This is much more readable and ultimately leads to more maintainable code.

## Usage

The matchers currently included in the `assert` package are:

1. IsNil

    ```cgo
    func TestIsNil(t *testing.T) {
        var s *string
        assert.With(t).
            That(s).
            IsNil()
    }
    ```

2. IsNotNil

    ```cgo
    func TestIsNotNil(t *testing.T) {
        var s string
        assert.With(t).
            That(s).
            IsNotNil()
    }
    ```

3. IsEqualTo

    ```cgo
    func TestEquals(t *testing.T) {
        got := Abs(-1)
        assert.With(t).
            That(got).
            IsEqualTo(1)
    }
    ```

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Metaleaf.io](http://github.com/metaleaf-io/)

## License

Copyright 2019 Metaleaf.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
