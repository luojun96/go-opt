package main

import (
	"encoding/binary"
	"io"
	"log"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 	exec()
	// 	r, err := os.Open("/hyperjun")
	// 	if err != nil {
	// 		log.Fatal("error openning 'a'\n")
	// 	}
	// 	defer Close(r)

	// 	r, err = os.Open("/hyperjun/jur")
	// 	if err != nil {
	// 		log.Fatal("error opening 'b'\n")
	// 	}
	// 	defer Close(r)
}

// Error Check Hell
type Point struct {
	Longitude     float64
	Latitude      float64
	Distance      float64
	ElevationGain float64
	ElevationLoss float64
}

func parse(r io.Reader) (*Point, error) {
	var p Point

	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parse1(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}
