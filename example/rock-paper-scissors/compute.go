package main

import "math/rand"

func compute() int {
	r := rand.Intn(3) + 1
	return r
}