package method

type User struct {
	ID         string
	Name       string
	Role       string
	Terminated bool
}

type matcher interface {
	Match(*User) bool
}

type UserRepository struct {
	users []*User
}

func (ur *UserRepository) Query(m matcher) ([]*User, error) { // want "^unexported interface 'matcher' used as parameter in exported method 'UserRepository.Query'$"
	var res []*User
	for _, u := range ur.users {
		if !m.Match(u) {
			continue
		}
		res = append(res, u)
	}
	return res, nil
}
