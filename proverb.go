// Package proverb provides literary proverbs as a library.
package proverb

import (
	_ "embed"
	"encoding/json"
	"math/rand"

	"github.com/pkg/errors"
)

// A Proverb is a wise sentence.
type Proverb struct {
	Author string
	Text   string
}

//go:embed proverb.json
var proverbBytes []byte

// A Collection is a collection of proverbs.
type Collection struct {
	Proverbs []Proverb
}

// NewCollection creates a collection of proverbs.
func NewCollection() (*Collection, error) {
	c := &Collection{}
	if err := json.Unmarshal(proverbBytes, &c.Proverbs); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return c, nil
}

// Sample samples a proverb.
func (c *Collection) Sample() Proverb {
	idx := rand.Intn(len(c.Proverbs))
	return c.Proverbs[idx]
}
