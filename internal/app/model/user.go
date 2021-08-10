package model

type User struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`
}

func (u *User) IsOverTwentyYearsOld() bool {
	return u.Age >= 20
}
