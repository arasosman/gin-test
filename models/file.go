package models

type File struct {
	ID             uint `json:"id" gorm:"primary_key"`
	Uuid           string
	Name           string `json:"name"`
	Path           string `json:"path"`
	FolderId       uint   `json:"folder_id"`
	InstallationId uint   `json:"installation_id"`
	CreatedAt      string `json:"created_at"`
}
