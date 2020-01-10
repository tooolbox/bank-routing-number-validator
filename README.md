## Summary

Validates Bank Routing Numbers

- Checks the first two digits against the allowed ranges
- Run the checksum algorithm on the last digit

See https://en.wikipedia.org/wiki/Routing_transit_number

## Install

```bash
$ go get github.com/tooolbox/bank-routing-number-validator
```

Or, more likely, import in your project.

## Usage

```golang
import "github.com/tooolbox/bank-routing-number-validator"

func main() {
	var err error
	err = validator.ABARoutingNumberIsValid("abcdabcde")  // error
	err = validator.ABARoutingNumberIsValid("1234567890") // error
	err = validator.ABARoutingNumberIsValid("021000021")  // success
	err = validator.ABARoutingNumberIsValid("21000021")   // success
}
```

## Credit
License for original JS version:
```
Copyright 2017 Chris Shaffer

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
