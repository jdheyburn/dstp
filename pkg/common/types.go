package common

type Output string

type Args []string

type Address string

type Port string

func (o Output) String() string {
	return string(o)
}

func (a Address) String() string {
	return string(a)
}

func (p Port) String() string {
	return string(p)
}
