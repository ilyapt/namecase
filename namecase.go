package namecase

import (
    "regexp"
    "strings"
)

// Precompile regular expressions once

var ufl = regexp.MustCompile(`\b(\w)`)
var ls = regexp.MustCompile(`(\'\w)\b`)

var mac = regexp.MustCompile(`\bMac[A-Za-z]{2,}[^aciozj]\b`)
var mc = regexp.MustCompile(`\bMc`)
var rmc = regexp.MustCompile(`\b(Ma?c)([A-Za-z]+)`)

var exp  = []string {
    "Macevicius",
    "Machado",
    "Machar",
    "Machin",
    "Machlin",
    "Macias",
    "Maciulis",
    "Mackie",
    "Mackle",
    "Macklin",
    "Macquarie",
    "Macomber",
    "Macin",
    "Mackintosh",
    "Macken",
    "Machen",
    "Machiel",
    "Maciol",
    "Mackell",
    "Macklem",
    "Mackrell",
    "Maclin",
    "Mackey",
    "Mackley",
    "Machell",
    "Machon",
}

var macmurdo = regexp.MustCompile(`\bMacmurdo`)
var macisaac = regexp.MustCompile(`\bMacisaac`)

var al = regexp.MustCompile(`\bAl(\s+\w)`)
var ap = regexp.MustCompile(`\bAp\b`)
var ben = regexp.MustCompile(`(\w\s+)Ben(\s+\w)`)
var dell = regexp.MustCompile(`\bDell([ae])\b`)
var daeiu = regexp.MustCompile(`\bD([aeiu])\b`)
var der = regexp.MustCompile(`\bDe([lr])\b`)
var el = regexp.MustCompile(`\bEl\b`)
var la = regexp.MustCompile(`\bLa\b`)
var le = regexp.MustCompile(`\bL([eo])\b`)
var van = regexp.MustCompile(`\bVan(\s+\w)`)
var von = regexp.MustCompile(`\bVon\b`)

func NameCase(name string) string {
    name = strings.ToLower(name)
    name = ufl.ReplaceAllStringFunc(name, func(s string) string {
        r := ufl.FindStringSubmatch(s)
        return strings.ToUpper(r[1])
	})
    name = ls.ReplaceAllStringFunc(name, func(s string) string {
        r := ls.FindStringSubmatch(s)
        return strings.ToLower(r[1])
	})

    if mac.MatchString(name) || mc.MatchString(name) {
        name = rmc.ReplaceAllStringFunc(name, func(s string) string {
            r := rmc.FindStringSubmatch(s)
            for _, ex := range exp {
                if r[0] == ex { return r[0] }
            }
            return r[1]+strings.ToUpper(r[2][:1])+r[2][1:]
        })
    }

    name = macmurdo.ReplaceAllString(name, "MacMurdo")
    name = macisaac.ReplaceAllString(name, "MacIsaac")

    name = al.ReplaceAllString(name, "al$1")
    name = ap.ReplaceAllString(name, "ap")
    name = ben.ReplaceAllString(name, "${1}ben$2")
    name = dell.ReplaceAllString(name, "dell$1")
    name = daeiu.ReplaceAllString(name, "d$1")
    name = der.ReplaceAllString(name, "de$1")
    name = el.ReplaceAllString(name, "el")
    name = la.ReplaceAllString(name, "la")
    name = le.ReplaceAllString(name, "l$1")
    name = van.ReplaceAllString(name, "van$1")
    name = von.ReplaceAllString(name, "von")

    return name
}
