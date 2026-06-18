package repository

const (
	qUpdateUserRole = `
		UPDATE users_roles
		SET role_id = $1
		WHERE user_id = $2
	`
	qUpdateTokenByUserID = `
		UPDATE users
		SET token_version = token_version + 1
		WHERE user_id = $1
	`
	qDeleteAssignRoles = `
		DELETE FROM users_roles
		WHERE user_id = $1
	`
	qAssignRoles = `
		INSERT INTO users_roles (user_id, role_id)
		VALUES
	`
)
