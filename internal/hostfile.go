package internal

import (
	"bufio"
	"log"
	"os"
)

type Hostfile struct {
	name string
}

func NewHostFile(file string) *Hostfile {
	log.Printf("Going to read {%s}", file)

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Failed to read file contents, reason: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hf, _ := ParseLine(scanner.Text())
		log.Println(hf)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	return &Hostfile{
		name: "Arjen",
	}
}
