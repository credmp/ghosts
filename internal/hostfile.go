package internal

type Hostfile struct {
	name string
}

func NewHostFile() *Hostfile {
	return &Hostfile{
		name: "Arjen",
	}
}
