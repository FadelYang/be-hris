package repository

const (
	qCreate = `
		INSERT INTO permissions (name, description)
		VALUES ($1, $2)
	`
	qCount = `
		SELECT COUNT(*)
		FROM permissions;
	`
	qGet = `
		SELECT * FROM permissions
		LIMIT $1 OFFSET $2
	`
	qGetByID = `
		SELECT * FROM permissions
		WHERE id = ($1)
	`
	qUpdateByID = `
		UPDATE permissions
		SET
			name = COALESCE($1, name)
			description = COALESCE($2, description)
		WHERE id = $2
	`
	qDeletebyID = `
		DELETE FROM permissions
		WHERE id = $1
	`
)
