package pref

import (
	"fmt"
	"os/exec"
)

// DataType is the data type of a Pref key
type DataType int

// List of DataTypes
const (
	String DataType = iota
	Data
	Int
	float
	Bool
	Date
	Array
	ArrayAdd
	Dict
	DictAdd
)

// Pref prepresents a macos key for the defaults command
type Pref struct {
	Domain    string
	Key       string
	ValueType DataType
	Value     interface{}
}

// New creates a new Pref object
func New(domain string, key string, valueType DataType, value interface{}) Pref {
	return Pref{domain, key, valueType, value}
}

// Read reads the value of the Pref key on the host computer
func (p Pref) Read() {
	out, err := exec.Command("defaults", "read", p.Domain, p.Key).Output()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		output := string(out[:])
		fmt.Println(output)
	}
}
