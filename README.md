sensou
======
Sensou/Sensō(戦争) means war in japanese. Sensou parses and writes warcraft 3 files.

Sensou gives further access to modification of warcraft 3 maps since the world editor doesn't limit some values read from .w3u files. One such example is the uhpm (unit health points) can exceed the 50k limit and instead hit the memory limit that's 16,777,215!

A list of hackable fields will follow:
uhpm

Library syntax - Parse
-------------

	wu, err := w3u.Open("unit-file.w3u")

Library syntax - Write
--------------

	wu := w3u.New()

	u, err := unit.NewUnit("hpea", "h000")
	if err != nil {
		// handle error
	}
	err = u.AddField("uhpm", 0xFFFFFF)
	if err != nil {
		// handle err
	}
	wu.AddUnit(u)

	err = wu.Write("/tmp/unit-output.w3u")
	if err != nil {
		// handle err
	}

Library syntax - Print
--------------

	wu, err := w3u.Open("unit-file.w3u")
	if err != nil {
		// handle errors
	}
	wu.PrintUnits()

Installation
------------
Use `go get github.com/karlek/sensou`

   go get github.com/karlek/sensou

Design Decisions
----------------
Should the w3u struct contain the number of units, separators and alike?
This will give the possibility to write the whole struct to a file.

API documentation
-----------------
http://go.pkgdoc.org/github.com/karlek/sensou

public domain
-------------
This code is hereby release this code into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
