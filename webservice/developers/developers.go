package developers

// Developer represents a developer
type Developer struct {
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Language string `json:"Language"`
	Floor    int    `json:"Floor"`
}

// Developers contains all registered developers
var Developers = map[int]Developer{
	1: {
		Name:     "Charlie",
		Age:      23,
		Language: "Go",
		Floor:    5,
	},
	2: {
		Name:     "Bill",
		Age:      23,
		Language: "Java",
		Floor:    5,
	},
	3: {
		Name:     "Kimbley",
		Age:      28,
		Language: "PHP",
		Floor:    2,
	},
	4: {
		Name:     "Roger",
		Age:      34,
		Language: "PHP",
		Floor:    2,
	},
	5: {
		Name:     "Alex",
		Age:      28,
		Language: "Go",
		Floor:    5,
	},
	6: {
		Name:     "Micheal",
		Age:      28,
		Language: "C++",
		Floor:    2,
	},
	7: {
		Name:     "Jill",
		Age:      28,
		Language: "Python",
		Floor:    2,
	},
}
