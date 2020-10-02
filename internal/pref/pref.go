package pref

import (
	"fmt"
	"log"
	"os/exec"
)

// Pref prepresents a defaults key
type Pref struct {
	Domain   string
	Key      string
	ValueRep string
	Value    string
}

// New creates a new Pref object
func New(domain string, key string, valueRep string, value string) Pref {
	return Pref{domain, key, valueRep, value}
}

// Read reads current the value of the Pref key on the host computer
func (p Pref) Read() string {
	out, err := exec.Command("defaults", "read", p.Domain, p.Key).Output()
	if err != nil {
		log.Fatal("error: Invalid command:")
	}
	return string(out[:])
}

// Write executes the given Pref
func (p Pref) Write() {
	out, err := exec.Command("defaults", "write", p.Domain, p.Key, "-"+p.ValueRep, p.Value).Output()
	if err != nil {
		log.Fatal("error: Invalid command:", err)
	}
	fmt.Println(string(out[:]))
}
