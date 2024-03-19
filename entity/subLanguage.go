package entity

type SubLanguage struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:text;not null"`
	MaterialID  uint       `json:"material_id"`
	Material    Material   `json:"material,omitempty" gorm:"foreignKey:MaterialID"`
	Questions   []Question `json:"questions,omitempty"`
	CourseID    uint       `json:"course_id"`
}
