// sensou test cases
package main

import "testing"
import "os"

import "github.com/karlek/sensou/unit"
import "github.com/karlek/sensou/w3u"

var sensouRoot = os.Getenv("GOPATH") + "/src/github.com/karlek/sensou/"

func TestWrite(t *testing.T) {
	wu := w3u.New()

	unit, err := unit.NewUnit("hpea", "h001")
	if err != nil {
		t.Fail()
	}

	err = unit.AddField("una1", 1)
	if err != nil {
		t.Fail()
	}

	wu.AddUnit(unit)

	err = wu.Write("/tmp/a.w3u")
	if err != nil {
		t.Fail()
	}
}

func TestOpen(t *testing.T) {
	wu, err := w3u.Open(sensouRoot + "files/01.w3u")
	if err != nil {
		t.Fail()
	}

	wu.PrintUnits()
	// Output: ufrd = <nil>
}
