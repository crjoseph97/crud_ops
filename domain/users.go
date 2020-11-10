package domain

// query "git.hifx.in/lens/querybuilder/redshift"

// User holds their profile data
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name,omitempty"`
	Department int    `json:"dept"`
	Mobile     int    `json:"mob"`
	Address    string `json:"addr"`
}

// UserBuilder holds their profile data
type UserBuilder struct {
	Name       string `json:"name" form:"name"`
	ID         string `json:"id" form:"id"`
	Mob        string `json:"mob" form:"mob"`
	Address    string `json:"addr" form:"addr"`
	Department string `json:"dept" form:"dept"`
}

// UserQueryBuilder creates the query needed list all users according to their dept
func UserQueryBuilder(qb UserBuilder) (string, error) {
	params := make([]interface{}, 0)
	query := `SELECT * FROM users WHERE dept in (?);`
	params = append(params, qb.Department)
	query = RemoveUnwantedSpaces(ReplacePositionalParamsInQuery(query, params...))
	return query, nil
}

// CreateQueryBuilder creates the query needed to add a new user
func CreateQueryBuilder(qb UserBuilder) (string, error) {
	params := make([]interface{}, 0)
	query := `INSERT INTO users VALUES (?,?,?,?,?)`
	params = append(params, qb.ID, qb.Name, qb.Department, qb.Mob, qb.Address)
	query = RemoveUnwantedSpaces(ReplacePositionalParamsInQuery(query, params...))
	return query, nil
}

// UpdateQueryBuilder creates the query needed to update data of a user id
func UpdateQueryBuilder(qb UserBuilder, id int) (string, error) {
	params := make([]interface{}, 0)
	query := `UPDATE users SET name= ?, dept= ?, mob= ?, addr= ? WHERE id= ?`
	params = append(params, qb.Name, qb.Department, qb.Mob, qb.Address, id)
	query = RemoveUnwantedSpaces(ReplacePositionalParamsInQuery(query, params...))
	return query, nil
}

// DeleteQueryBuilder creates the query needed to delete data of a user id
func DeleteQueryBuilder(id int) (string, error) {
	params := make([]interface{}, 0)
	query := `DELETE from users WHERE id= ?`
	params = append(params, id)
	query = RemoveUnwantedSpaces(ReplacePositionalParamsInQuery(query, params...))
	return query, nil
}
