package main

var head = []*[4]byte{
	&[4]byte{'P', 'N', 'G', ' '},
	&[4]byte{'G', 'I', 'F', ' '},
	&[4]byte{'J', 'P', 'E', 'G'},
}

var heads = []*[4]byte{
	{'P', 'N', 'G', ' '},
	{'G', 'I', 'F', ' '},
	{'J', 'P', 'E', 'G'},
}

type language struct {
	name string
	age  int
}

var _ = [...]language{
	language{"C", 1972},
	language{"Python", 1991},
	language{"Go", 2009},
}

var _ = [...]language{
	{"C", 1972},
	{"Python", 1991},
	{"Go", 2009},
}

type LangCategory struct {
	dynamic bool
	strong  bool
}

var _ = map[LangCategory]map[string]int{
	{true, false}: {
		"Python": 1991,
		"Erlang": 1986,
	},
	{true, false}: {
		"JavaSctript": 1995,
	},
	{false, true}: {
		"Go":   2009,
		"Rust": 2010,
	},
	{false, false}: {
		"C": 1972,
	},
}
