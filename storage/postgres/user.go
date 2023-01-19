package postgres

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/authorization"
	"errors"
	"time"
)

// *=========================================================================
func (stg Postgres) AddNewUser(id string, req *authorization.CreateUserRequest) error {
	_, err := stg.homeDB.Exec(`INSERT INTO "user" 
	(
		"id", 
		"fname", 
		"lname", 
		"username", 
		"password",
		"user_type", 
		"address", 
		"phone", 
		"created_at"
	) VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5, 
		$6, 
		$7, 
		$8, 
		now()
	)`, id, req.Fname, req.Lname, req.Username, req.Password, req.UserType, req.Address, req.Phone)
	if err != nil {
		return err
	}
	return nil
}

// *=========================================================================
func (stg Postgres) GetUserById(id string) (*authorization.User, error) {
	res := &authorization.User{}
	var deletedAt *time.Time
	var updatedAt *string

	err := stg.homeDB.QueryRow(`SELECT 
		"id", 
		"fname", 
		"lname", 
		"username", 
		"password",
		"user_type", 
		"address", 
		"phone", 
		"created_at",
		"updated_at", 
		"deleted_at"
    FROM "user" WHERE id = $1`, id).Scan(
		&res.Id, 
		&res.Fname, 
		&res.Lname, 
		&res.Username, 
		&res.Password, 
		&res.UserType,
		&res.Address, 
		&res.Phone, 
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
func (stg Postgres) GetUserList(offset, limit int, search string) (*authorization.GetUserListResponse, error) {
	res := &authorization.GetUserListResponse{
		Users: make([]*authorization.User, 0),
	}
	rows, err := stg.homeDB.Queryx(`SELECT
		"id", 
		"fname", 
		"lname", 
		"username", 
		"password",
		"user_type", 
		"address", 
		"phone",
		"created_at",
		"updated_at"
	FROM "user" WHERE deleted_at IS NULL AND ("username" || "fname" || "lname" ILIKE '%' || $1 || '%')
		LIMIT $2
		OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		a := &authorization.User{}

		var updatedAt *string

		err := rows.Scan(
			&a.Id, 
			&a.Fname, 
			&a.Lname, 
			&a.Username, 
			&a.Password, 
			&a.UserType,
			&a.Address, 
			&a.Phone, 
			&a.CreatedAt, 
			&updatedAt,
		)
		if err != nil {
			return res, err
		}

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		res.Users = append(res.Users, a)
	}
	return res, err
}

// *=========================================================================
func (stg Postgres) UpdateUser(box *authorization.UpdateUserRequest) error {
	res, err := stg.homeDB.NamedExec(
	`UPDATE "user"  
		SET 
			"fname"=:f, 
			"lname"=:l, 
			"username"=:u, 
			"password"=:p, 
			"user_type"=:ut, 
			"address"=:a, 
			"phone"=:ph, 
			"updated_at"=now() 
		WHERE deleted_at IS NULL AND id=:id`, map[string]interface{}{
			"id": box.Id,
			"f":  box.Fname, 
			"l": box.Lname, 
			"u": box.Username, 
			"p": box.Password,
			"ut": box.UserType, 
			"a": box.Address, 
			"ph": box.Phone,
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
func (stg Postgres) GetUserByUsername(username string) (*authorization.User, error) {
	res := &authorization.User{}
	var deletedAt *time.Time
	var updatedAt *string

	err := stg.homeDB.QueryRow(`SELECT 
		"id", 
		"fname", 
		"lname", 
		"username", 
		"password",
		"user_type", 
		"address", 
		"phone", 
		"created_at",
		"updated_at", 
		"deleted_at"
    FROM "user" WHERE username = $1`,username).Scan(
		&res.Id, 
		&res.Fname, 
		&res.Lname, 
		&res.Username, 
		&res.Password, 
		&res.UserType,
		&res.Address, 
		&res.Phone, 
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
