# NameCase

NameCase - Go package to fix the case of people's names, based on the
Perl [Lingua-EN-NameCase](http://cpansearch.perl.org/src/SUMMER/Lingua-EN-NameCase-1.15/) module.

## Examples

```
package main

import (
    "fmt"
    "github.com/ilyapt/namecase"
)

function main() {
    name := "GEORGE WASHINGTON"
    name = namecase.NameCase(name)
    fmt.Println(name)
}
```

## Description

Forenames and surnames are often stored either wholly in UPPERCASE
or wholly in lowercase. This module allows you to convert names into
the correct case where possible.

Although forenames and surnames are normally stored separately if they
do appear in a single string, whitespace separated, NameCase deal
correctly with them.

NameCase currently correctly name cases names which include any of the
following:
    Mc, Mac, al, el, ap, da, de, delle, della, di, du, del, der,
    la, le, lo, van and von.

It correctly deals with names which contain apostrophies and hyphens too.