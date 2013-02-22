//
package main

import "fmt"
import "log"

import "github.com/karlek/sensou/unit"
import "github.com/karlek/sensou/w3u"

func main() {
	err := play()
	if err != nil {
		log.Fatalln(err)
	}
}

func play() (err error) {
	w3 := w3u.New()

	u, err := unit.NewUnit("hpea", "h001")
	if err != nil {
		return err
	}

	err = u.AddField("una1", 1)
	if err != nil {
		return err
	}

	w3.AddUnit(u)

	w3.PrintUnits()
	w3.Write("a.w3u")

	w3 = w3u.New()
	w3, err = w3u.Open("a.w3u")
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", w3)

	return nil
}
