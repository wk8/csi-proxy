package apiversions

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

// TODO wkpo move to internal?

// TODO wkpo unit tests...

// TODO wkpo have a special Dev version that's the latest one?
// TODO wkpo? keep? if so, make it the latest version!
// TODO wkpo or master?
var Dev = Version{}

// API version names follow k8s style version names:
// valid API version names examples: "v1", "v1-alpha2", "v2-beta3", etc...
var nameRegex = regexp.MustCompile("^v([0-9]*)(?:(alpha|beta)([1-9][0-9]*))?$")

type qualifier uint

const (
	stable qualifier = iota
	beta
	alpha
)

type Version struct {
	// major version number, eg 1 for "v1", 2 for "v2-beta3"
	major uint

	// qualifier is "alpha" or "beta"
	qualifier qualifier

	// the qualifier's version, eg 2 for "v1-alpha2" or 3 for "v2-beta3"
	qualifierVersion uint

	rawName string
}

// NewVersion will panic if passed an invalid version name.
func NewVersion(name string) Version {
	matches := nameRegex.FindStringSubmatch(name)
	if len(matches) < 2 {
		panic(fmt.Errorf("invalid version name: %q", name))
	}

	major, err := toInt(matches[1], name)
	if err != nil {
		panic(err)
	}

	var (
		qualifier        qualifier
		qualifierVersion uint
	)
	if len(matches) >= 4 {
		switch matches[2] {
		case "alpha":
			qualifier = alpha
		case "beta":
			qualifier = beta
		}

		qualifierVersion, err = toInt(matches[3], name)
		if err != nil {
			panic(err)
		}
	}

	return Version{
		major:            major,
		qualifier:        qualifier,
		qualifierVersion: qualifierVersion,
		rawName:          name,
	}
}

func toInt(s, name string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, errors.Wrapf(err, "unable to convert %q from version name %q to int", s, name)
	}
	return uint(i), nil
}

type Comparison int

const (
	Lesser  Comparison = -1
	Equal   Comparison = 0
	Greater Comparison = 1
)

// Compare return Lesser if v < other (ie other is more recent), Equal if v == other,
// and Greater if v > other (ie v is more recent)
func (v Version) Compare(other Version) Comparison {
	if cmp := compareUints(v.major, other.major); cmp != Equal {
		return cmp
	}
	if cmp := compareUints(uint(v.qualifier), uint(other.qualifier)); cmp != Equal {
		return cmp
	}
	return compareUints(v.qualifierVersion, other.qualifierVersion)
}

func compareUints(a, b uint) Comparison {
	if a < b {
		return Lesser
	}
	if a > b {
		return Greater
	}
	return Equal
}

func (v Version) String() string {
	return v.rawName
}
