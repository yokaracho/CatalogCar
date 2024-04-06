package db

const (
	insertRegNum = `INSERT INTO cars (reg_num) VALUES ($1)`

	deleteCar = `DELETE FROM cars WHERE "id" = $1`

	updateCar = `UPDATE cars
				SET reg_num = $2,
					mark = $3,
					model = $4,
					year = $5
				WHERE id = $1;`
)
