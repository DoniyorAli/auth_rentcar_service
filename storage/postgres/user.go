package postgres

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/blogpost"
	"errors"
	"time"
)

// *=========================================================================
func (stg Postgres) AddNewUser(id string, box *blogpost.CreateUserRequest) error {
	_, err := stg.homeDB.Exec(`INSERT INTO "user" 
	(
		id,
		username,
		password,
		user_type
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		id,
		box.Username,
		box.Password,
		box.UserType,
	)
	if err != nil {
		return err
	}
	return nil
}

// *=========================================================================
func (stg Postgres) GetUserById(id string) (*blogpost.User, error) {
	res := &blogpost.User{}
	var deletedAt *time.Time
	var updatedAt *string

	err := stg.homeDB.QueryRow(`SELECT 
		id,
		username,
		password,
		user_type,
		created_at,
		updated_at,
		deleted_at
    FROM "user" WHERE id = $1`, id).Scan(
		&res.Id,
		&res.Username,
		&res.Password,
		&res.UserType,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if deletedAt != nil {
		return res, errors.New("user not found")
	}

	return res, err
}

// *=========================================================================
func (stg Postgres) GetUserList(offset, limit int, search string) (*blogpost.GetUserListResponse, error) {
	res := &blogpost.GetUserListResponse{
		User: make([]*blogpost.User, 0),
	}
	rows, err := stg.homeDB.Queryx(`SELECT
	id,
	username,
	password,
	user_type,
	created_at,
	updated_at
	FROM "user" WHERE deleted_at IS NULL AND (username ILIKE '%' || $1 || '%')
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		a := &blogpost.User{}

		var updatedAt *string

		err := rows.Scan(
			&a.Id,
			&a.Username,
			&a.Password,
			&a.UserType,
			&a.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return res, err
		}

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		res.User = append(res.User, a)
	}
	return res, err
}

// *=========================================================================
func (stg Postgres) UpdateUser(box *blogpost.UpdateUserRequest) error {
	res, err := stg.homeDB.NamedExec(`UPDATE "user"  SET password=:p, updated_at=now() WHERE deleted_at IS NULL AND id=:id`, map[string]interface{}{
		"id": box.Id,
		"p":  box.Password,
	})
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("user not found")
}

// *=========================================================================
func (stg Postgres) DeleteUser(id string) error {
	res, err := stg.homeDB.Exec(`UPDATE "user" SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect > 0 {
		return nil
	}
	return errors.New("user not found")
}

// *=========================================================================
func (stg Postgres) GetUserByUsername(username string) (*blogpost.User, error) {
	res := &blogpost.User{}
	var deletedAt *time.Time
	var updatedAt *string

	err := stg.homeDB.QueryRow(`SELECT 
		id,
		username,
		password,
		user_type,
		created_at,
		updated_at,
		deleted_at
    FROM "user" WHERE username = $1`, username).Scan(
		&res.Id,
		&res.Username,
		&res.Password,
		&res.UserType,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if deletedAt != nil {
		return res, errors.New("user not found")
	}

	return res, err
}
