package main

type Entry struct {
	Title   string
	Covers  []int
	Key     string
	Authors struct {
		Type   struct{ Key string }
		Author struct{ Key string }
	}
	Type           struct{ Key string }
	Subjects       []string
	LatestRevision int
	Created        struct {
		Type  string
		Value string
	}
	LastModified struct {
		Type  string
		Value string
	}
}
