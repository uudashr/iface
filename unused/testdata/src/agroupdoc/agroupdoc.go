package user

import "database/sql"

type (
	User struct {
		ID   string
		Name string
	}

	// UserRepository provides access to user data.
	UserRepository interface { // want "interface 'UserRepository' is declared but not used within the package"
		UserOf(id string) (*User, error)
	}

	UserRepositorySQL struct {
		DB *sql.DB
	}
)

func (r *UserRepositorySQL) UserOf(id string) (*User, error) {
	var u User
	err := r.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", u.ID).Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}

	return &u, nil
}