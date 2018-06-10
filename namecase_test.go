package namecase

import (
    "strings"
    "testing"
    "github.com/magiconair/properties/assert"
)

var names = []string{
    "Keith", "Leigh-Williams", "McCarthy",
    "O'Callaghan", "St. John", "von Streit",
    "van Dyke", "Van", "ap Llwyd Dafydd",
    "al Fahd", "Al", "el Grecco",
    "David ben Gurion", "Ben-Gurion", "Ben",
    "da Vinci",
    "di Caprio", "du Pont", "de Legate",
    "del Crond", "der Sind", "van der Post",
    "von Trapp", "la Poisson", "le Figaro",
    "Mack Knife", "Dougal MacDonald",
    // Mac exceptions
    "Machin", "Machlin", "Machar",
    "Mackle", "Macklin", "Mackie",
    "Macquarie", "Machado", "Macevicius",
    "Maciulis", "Macias", "MacMurdo",
    "Mackrell", "Maclin", "McConnachie",
}

var lowerNames = func() []string {
    var out []string
    for _, name := range names { out = append(out, strings.ToLower(name)) }
    return out
}()

func Test_NameCase(t *testing.T) {
    for i, name := range names {
        assert.Equal(t, name, NameCase(lowerNames[i]))
    }
}