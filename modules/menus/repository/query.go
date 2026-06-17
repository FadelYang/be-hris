package repository

const (
	qCreate = `
		INSERT INTO menus (name, slug, parent_menu_id)
		VALUES ($1, $2, $3)
	`
	qCount = `
		SELECT COUNT(*)
		FROM menus;
	`
	qGet = `
		SELECT * FROM menus
		LIMIT $1 OFFSET $2
	`
	qGetByID = `
		SELECT * FROM menus
		WHERE id = ($1)
	`
	qUpdateByID = `
		UPDATE menus
		SET
			name = COALESCE($1, name),
			slug = COALESCE($2, slug),
			parent_menu_id = COALESCE($3, parent_menu_id),
		WHERE id = $4
	`
	qDeletebyID = `
		DELETE FROM menus
		WHERE id = $1
	`
)
