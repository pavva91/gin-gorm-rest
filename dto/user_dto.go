package dto

import "github.com/pavva91/gin-gorm-rest/models"

type UserDTO struct {
	ID       uint   `json:"userID"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (dto *UserDTO) ToModel() *models.Users {
	var model *models.Users
	model.Name = dto.Name
	model.Surname = dto.Surname
	model.Username = dto.Username
	model.Email = dto.Email
	return model
}

func (dto *UserDTO) ToDTO(model models.Users) {
	dto.ID = model.ID
	dto.Name = model.Name
	dto.Surname = model.Surname
	dto.Username = model.Username
	dto.Email = model.Email
}

func (dto *UserDTO) ToDTOs(models []models.Users) (dtos []UserDTO) {
	dtos = make([]UserDTO, len(models))
	for i, v := range models {
		dtos[i].ToDTO(v)
	}
	return dtos
}
