package repository

import (
	"HR/internal/domain"
	"HR/internal/infrastructure"
)

func CreateEmployee(emp domain.Employee) error {
	return infrastructure.DB.Create(&emp).Error
}

func GetAllEmployees() ([]domain.Employee, error) {
	var employees []domain.Employee
	result := infrastructure.DB.Find(&employees)
	return employees, result.Error
}

func GetEmployeeByID(id string) (domain.Employee, error) {
	var emp domain.Employee
	result := infrastructure.DB.First(&emp, "id = ?", id)
	return emp, result.Error
}

func UpdateEmployee(id string, emp domain.Employee) error {
	emp.ID = id
	return infrastructure.DB.Save(&emp).Error
}

func DeleteEmployee(id string) error {
	return infrastructure.DB.Unscoped().Delete(&domain.Employee{}, "id = ?", id).Error
}

func SoftDeleteEmployee(id string) error {
	return infrastructure.DB.Delete(&domain.Employee{}, "id = ?", id).Error
}

func RecoverEmployee(id string) error {
	return infrastructure.DB.Unscoped().Model(&domain.Employee{}).
		Where("id = ?", id).Update("deleted_at", nil).Error
}
