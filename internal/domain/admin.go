package domain

import "time"

type Admin struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
