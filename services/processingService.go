package service

import (
	model "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/models"
	repository "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/repositories"
)

func NewProcessing(processing model.Processing) error {
	return repository.ProcessingSave(processing)
}

func ListAllProcessing() ([]model.Processing, error) {
	return repository.ProcessingGetAll()
}
