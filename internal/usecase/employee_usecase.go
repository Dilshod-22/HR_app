package usecase

import (
	"HR/internal/domain"
	"HR/internal/repository"
)

func GetEmployee() ([]domain.Employee, error) {
	return repository.GetAllEmployees()
}

func CreateEmployee(emp domain.Employee) error {
	return repository.CreateEmployee(emp)
}

func UpdateEmployee(id uint, emp domain.Employee) error {
	return repository.UpdateEmployee(id, emp)
}

func DeleteEmployee(id uint) error {
	return repository.DeleteEmployee(id)
}

func SoftDeleteEmployee(id uint) error {
	return repository.SoftDeleteEmployee(id)
}

func RecoverEmployee(id uint) error {
	return repository.RecoverEmployee(id)
}
