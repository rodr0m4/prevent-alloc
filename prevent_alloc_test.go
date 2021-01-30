package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"testing"
)

type Subject struct {}

func (s Subject) ShouldEscape() bool {
	n, err := rand.Int(rand.Reader, big.NewInt(10))

	if err != nil {
		log.Panic(err)
	}

	// 2 / 10 possibilities
	return n.Int64() >= 8
}

func (s Subject) IsZero() bool {
	return Subject{} == s
}

func Escaping() *Subject {
	var s Subject

	if s.ShouldEscape() {
		return &s
	}

	return nil
}

func Test_That_Always_Passes(t *testing.T) {
	e := Escaping()
	fmt.Printf("e is %v\n", e)
	if e != nil && !e.IsZero() {
		t.Errorf("not nil nor zero subject")
	}
}
