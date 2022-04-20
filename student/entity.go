package student

import "time"

type Student struct {
	ID        uint
	Name      string
	Major     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
