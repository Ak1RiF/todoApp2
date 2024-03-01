package models

type User struct {
	Id                    int
	Username              string
	PasswordHash          string
	AvatarUrl             string
	TotalExperience       int
	AmountExperienceToLvl int
	Lvl                   int
}
