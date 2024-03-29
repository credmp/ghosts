package internal

import (
	"errors"
	"log"
	"net"
	"regexp"
	"strings"
)

// Host entry structure that resembles a line in a hosts file
type HostEntry struct {
	Name     string
	IP       net.IP
	Aliasses []string
	Comment  string
}

// Create a new HostEntry with the name and ip address
func NewHostEntry(name string, ip net.IP) *HostEntry {
	return &HostEntry{
		Name: name,
		IP:   ip,
	}
}

func (he *HostEntry) CanResolve(name string) bool {
	if len(he.Name) == 0 {
		return false
	}
	if strings.HasSuffix(name, he.Name) {
		return true
	}

	return false
}

func (he *HostEntry) WouldResolve(name string) bool {
	if len(name) == 0 || len(he.Name) == 0 {
		return false
	}

	if strings.HasSuffix(he.Name, name) {
		return true
	}

	return false
}

func (he *HostEntry) SwapInName(name string) {
	oldname := he.Name
	he.Name = name
	he.Aliasses = append(he.Aliasses, oldname)
}

func (he *HostEntry) ReplaceIP(ip string) {
	new_ip := net.ParseIP(ip)
	if new_ip == nil {
		log.Fatal("Failed to convert ip address", ip)
	}
	he.IP = new_ip
}

func (he *HostEntry) AddAlias(name string) {
	for _, e := range he.Aliasses {
		if e == name {
			return // already in the list
		}
	}

	he.Aliasses = append(he.Aliasses, name)
}

// Parse a string into a HostEntry
func ParseLine(line string) (*HostEntry, error) {
	comment := regexp.MustCompile(`^#\s*(?P<comment>.+)\s*$`)
	entry := regexp.MustCompile(`^(?P<ip>.+?)\s+(?P<name>.+?)(\s+(?P<aliasses>[^#]+))?(#\s*(?P<comment>.*))?$`)

	if len(line) == 0 {
		return &HostEntry{}, nil
	}

	if comment.MatchString(line) {

		result := ApplyRegexp(comment, line)
		return &HostEntry{Comment: result["comment"]}, nil
	}

	result := ApplyRegexp(entry, line)

	ip := net.ParseIP(result["ip"])
	if ip == nil {
		log.Fatal("Could not parse IP - TODO: deal with this nicer")
	}

	var aliasses []string
	if len(result["aliasses"]) > 0 {
		re := regexp.MustCompile("[\\s\\t]")
		split := re.Split(result["aliasses"], -1)
		for i := range split {
			s := strings.TrimSpace(split[i])
			if len(s) > 0 {
				aliasses = append(aliasses, s)
			}
		}
	}

	he := HostEntry{
		IP:       ip,
		Name:     result["name"],
		Comment:  result["comment"],
		Aliasses: aliasses,
	}
	return &he, errors.New("Could not parse line")
}

func ApplyRegexp(re *regexp.Regexp, line string) map[string]string {

	// Find match and capture groups
	match := re.FindStringSubmatch(line)

	// Create a map to hold the named capture groups
	groupNames := re.SubexpNames()
	result := make(map[string]string)
	for i, name := range groupNames {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
