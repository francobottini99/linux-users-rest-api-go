package repository

import (
	"fmt"
	"log"

	model "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/models"
	"github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/repositories/database"
)

func ProcessingSave(process model.Processing) error {
	stm, err := database.ProcessingGetDBConnection().Prepare("INSERT INTO processing (process, free_memory, swap) VALUES (?, ?, ?);")

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Failed save processing: '%v'", err)
	}

	defer stm.Close()

	_, err = stm.Exec(process.Process, process.FreeMemory, process.Swap)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Failed save processing: '%v'", err)
	}

	return nil
}

func ProcessingGetAll() ([]model.Processing, error) {
	var result []model.Processing

	stm, err := database.ProcessingGetDBConnection().Query("SELECT * FROM processing;")

	if err != nil {
		log.Println(err)
		return result, fmt.Errorf("Failed get all processing: '%v'", err)
	}

	defer stm.Close()

	for stm.Next() {
		var process model.Processing

		err = stm.Scan(&process.Id, &process.Process, &process.FreeMemory, &process.Swap)

		if err != nil {
			log.Println(err)
			return result, fmt.Errorf("Failed get all processing: '%v'", err)
		}

		result = append(result, process)
	}

	return result, nil
}
