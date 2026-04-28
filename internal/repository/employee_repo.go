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

func GetEmployeeByID(id uint) (domain.Employee, error) {
	var emp domain.Employee
	result := infrastructure.DB.First(&emp, id)
	return emp, result.Error
}

func UpdateEmployee(id uint, emp domain.Employee) error {
	emp.ID = id
	return infrastructure.DB.Save(&emp).Error
}

// Hard delete — bazadan butunlay o'chiradi
func DeleteEmployee(id uint) error {
	return infrastructure.DB.Unscoped().Delete(&domain.Employee{}, id).Error
}

// Soft delete — deleted_at ni belgilaydi
func SoftDeleteEmployee(id uint) error {
	return infrastructure.DB.Delete(&domain.Employee{}, id).Error
}

// Soft delete ni bekor qiladi
func RecoverEmployee(id uint) error {
	return infrastructure.DB.Unscoped().Model(&domain.Employee{}).
		Where("id = ?", id).Update("deleted_at", nil).Error
}
