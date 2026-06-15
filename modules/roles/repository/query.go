package repository

const (
	qCreate = `
		INSERT INTO roles (name)
		VALUES ($1)
	`
	qGet = `
		SELECT * FROM roles
	`
	qGetByID = `
		SELECT * FRIM roles
		WHERE id = ($1)
	`
	qUpdateByID = `
		UPDATE roles
		SET
			name = COALESCE($1, name)
		WHERE id = $2
	`
	qDeletebyID = `
		DELETE FROM roles
		WHERE id = $1
	`
)
