package validator

import "../dto"

func ValidateEmployee(p_Employee dto.Employee) bool {
	if p_Employee.Code != "" {
		if p_Employee.Name != "" {
			if p_Employee.Address != "" {
				if p_Employee.Salary != 0 {
					return true
				}
			}
		}
	}

	return false
}
