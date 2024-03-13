package dtos

type InputUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OutputQuestDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Dificulty   string `json:"dificulty"`
	Completed   bool   `json:"completed"`
}
