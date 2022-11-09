package main

type LAYOUT string

/*
    The format of the layout remains the same throughout this program, although
	it could be setup to have different setup.

	For 2x4 ortho (grid) format, the keys for left hand are stored from right to
	left, top to bottom.  I.e., the keys for right hand are stored from left to
	right, top to bottom.
*/

const artsey LAYOUT = "artseyio"

var combos = []int{
	0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80,
	0x03, 0x06, 0xc0, 0x05, 0x0a, 0x09,
	0x30, 0x60, 0xc0, 0x50, 0xa0, 0x90,
}

func New() (k LAYOUT, error int) {
	return k, 0
}
