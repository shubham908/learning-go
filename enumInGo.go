package main

const (
	UNKNOWN = 0
	FEMALE = 1
	MALE = 2
)

// Using iota will make go assign value incremently
const (
	ZERO = iota
	ONE = iota
	TWO = iota
)

type Gender int 
const (
	UNKNOWN Gender = iota
	FEMALE
	MALE
)