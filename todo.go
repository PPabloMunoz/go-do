package main

import "time"

type todo struct {
	Name string     `json:"name"`
	Done *time.Time `json:"done"`
}

func (t *todo) Toggle() {
	if t.Done != nil {
		t.Done = nil
	} else {
		now := time.Now()
		t.Done = &now
	}
}
