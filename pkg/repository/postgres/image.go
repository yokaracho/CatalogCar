package db

import (
	"CatalogCar/pkg/model"
	"context"
	"strconv"
)

func (r *Repository) GetCars(ctx context.Context, filters model.CarFilter) ([]*model.CarModel, error) {
	var cars []*model.CarModel
	var args []interface{}
	query := `
		SELECT 
			c.id, 
			CASE WHEN c.reg_num IS NULL THEN '' ELSE c.reg_num END AS reg_num,
			CASE WHEN c.mark IS NULL THEN '' ELSE c.mark END AS mark,
			CASE WHEN c.model IS NULL THEN '' ELSE c.model END AS model,
 			CASE WHEN c.year IS NULL THEN 0 ELSE c.year END AS year,
   			CASE WHEN o.id IS NULL THEN 0 ELSE o.id END AS owner_id,
			CASE WHEN o.name IS NULL THEN '' ELSE o.name END AS name,
			CASE WHEN o.surname IS NULL THEN '' ELSE o.surname END AS surname,
			CASE WHEN o.patronymic IS NULL THEN '' ELSE o.patronymic END AS patronymic
		FROM 
			cars c
		LEFT JOIN 
			owners o ON c.owner_id = o.id
		WHERE 
			1=1
	`

	if filters.RegNum != "" {
		query += " AND c.reg_num = $" + strconv.Itoa(len(args)+1)
		args = append(args, filters.RegNum)
	}

	if filters.Mark != "" {
		query += " AND c.mark = $" + strconv.Itoa(len(args)+1)
		args = append(args, filters.Mark)
	}

	if filters.Model != "" {
		query += " AND c.model = $" + strconv.Itoa(len(args)+1)
		args = append(args, filters.Model)
	}

	if filters.Year != 0 {
		query += " AND c.year = $" + strconv.Itoa(len(args)+1)
		args = append(args, filters.Year)
	}

	query += " ORDER BY c.id OFFSET $" + strconv.Itoa(len(args)+1) + " LIMIT $" + strconv.Itoa(len(args)+2)

	args = append(args, filters.Offset, filters.Limit)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.CarModel
		if err := rows.Scan(&car.CarID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.OwnerID, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic); err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *Repository) InsertCars(ctx context.Context, regNums []string) (int64, error) {
	var totalRowsAffected int64

	for _, regNum := range regNums {
		result, err := r.pool.Exec(ctx, insertRegNum, regNum)
		if err != nil {
			return totalRowsAffected, err
		}
		totalRowsAffected += result.RowsAffected()
	}

	return totalRowsAffected, nil
}

func (r *Repository) UpdateInfo(ctx context.Context, car *model.CarModel) error {
	_, err := r.pool.Exec(ctx, updateCar, car.CarID, car.RegNum, car.Mark, car.Model, car.Year)
	if err != nil {
		return err
	}
	return err
}

func (r *Repository) UpdateOwner(ctx context.Context, owner *model.PeopleModel) error {
	_, err := r.pool.Exec(ctx, updateOwner, owner.OwnerID, owner.Name, owner.Surname, owner.Patronymic)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteCarByID(ctx context.Context, id int) (int64, error) {
	result, err := r.pool.Exec(ctx, deleteCar, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), err
}
