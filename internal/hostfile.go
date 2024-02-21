package internal

import (
	"bufio"
	"log"
	"os"
)

type Hostfile struct {
	Entries []*HostEntry
}

func NewHostFile(file string) *Hostfile {
	log.Printf("Going to read {%s}", file)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Failed to read file contents, reason: ", err)
	}
	defer f.Close()
	hf := Hostfile{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		he, _ := ParseLine(scanner.Text())
		hf.Entries = append(hf.Entries, he)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	return &hf
}
