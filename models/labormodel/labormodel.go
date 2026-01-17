package labormodel

import (
	"go-web/config"
	"go-web/entities"
)

func GetAll()[]entities.Labor {
	rows, err := config.DB.Query("SELECT * FROM labor")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var labors []entities.Labor

	for rows.Next(){
		var labor entities.Labor
		if err := rows.Scan(&labor.Id, &labor.Nama_lab, &labor.Created_at, &labor.Updated_at); err != nil {
			panic(err)
		}
		labors = append(labors, labor)
	}
	return labors
}