package unit

const (
	FIND_ALL = `
		SELECT 
			u.id, 
			u.name,
			u.type,
			u.status,
			u."lastUpdated"
		FROM 
			unit u
		WHERE
			u.id IS NOT NULL
	`

	FIND_BY_ID = `
		SELECT  
			u.id, 
			u.name,
			u.type,
			u.status,
			u."lastUpdated" 
		FROM 
			unit u
		where u.id = $1
	`

	FIND_BY_NAME = `
		SELECT  
			u.id, 
			u.name,
			u.type,
			u.status,
			u."lastUpdated"  
		FROM 
			unit u
		where u.name = $1
	`

	INSERT = `
		INSERT INTO unit (
			id, 
			name, 
			type, 
			status, 
			"lastUpdated"
		) 
		VALUES (
			$1, 
			$2, 
			$3, 
			$4, 
			now()
		)
	`

	UPDATE_BY_ID = `
		UPDATE unit
		SET 
			name = $2,
			type = $3,
			status = $4,
			"lastUpdated" = NOW()
		WHERE 
			id = $1
	`
)
