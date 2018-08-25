package user

type Model struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Create(name string, email string) *Model {
	if name != "" && email != "" {
		return &Model{name, email}
	}
	return nil
}
