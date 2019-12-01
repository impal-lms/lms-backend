package repository

import "github.com/impal-lms/lms-backend/domain"

func (g *GORM) GetAllClassroom() ([]domain.Classroom, error) {
	var classroom []domain.Classroom
	if err := g.DB.Model(&classroom).Find(&classroom).Error; err != nil {
		return classroom, err
	}

	return classroom, nil
}

func (g *GORM) CreateClassroom(classroom domain.Classroom) (domain.Classroom, error) {
	if err := g.DB.Model(&classroom).Create(&classroom).Error; err != nil {
		return classroom, err
	}

	return classroom, nil
}

func (g *GORM) GetClassroomByID(id int64) (domain.Classroom, error) {
	var classroom domain.Classroom
	if err := g.DB.Model(&classroom).Where("id = ?", id).Find(&classroom).Error; err != nil {
		return classroom, err
	}

	return classroom, nil
}

func (g *GORM) UpdateClassroom(classroom domain.Classroom) error {
	if err := g.DB.Model(&classroom).Where("id = ?", classroom.ID).Update(&classroom).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteClassroomByID(id int64) error {
	var classroom domain.Classroom
	if err := g.DB.Model(&classroom).Where("id = ?", id).Delete(&classroom).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) GetAllStudentClassroom(studentID, classroomID int64) ([]domain.StudentClassroom, error) {
	var studentClassroom []domain.StudentClassroom

	if studentID != 0 {
		if classroomID != 0 {
			if err := g.DB.Model(&studentClassroom).Where("student_id = ? AND classroom_id = ?", studentID, classroomID).Find(&studentClassroom).Error; err != nil {
				return studentClassroom, err
			}
		}

		if err := g.DB.Model(&studentClassroom).Where("student_id = ?", studentID).Find(&studentClassroom).Error; err != nil {
			return studentClassroom, err
		}
	}

	if classroomID != 0 {
		if err := g.DB.Model(&studentClassroom).Where("classroom_id = ?", classroomID).Find(&studentClassroom).Error; err != nil {
			return studentClassroom, err
		}
	}

	if err := g.DB.Model(&studentClassroom).Find(&studentClassroom).Error; err != nil {
		return studentClassroom, err
	}

	return studentClassroom, nil
}

func (g *GORM) CreateStudentClassroom(studentClassroom domain.StudentClassroom) (domain.StudentClassroom, error) {
	if err := g.DB.Model(&studentClassroom).Create(&studentClassroom).Error; err != nil {
		return studentClassroom, err
	}

	return studentClassroom, nil
}

func (g *GORM) GetStudentClassroomByID(id int64) (domain.StudentClassroom, error) {
	var studentClassroom domain.StudentClassroom
	if err := g.DB.Model(&studentClassroom).Where("id = ?", id).Find(&studentClassroom).Error; err != nil {
		return studentClassroom, err
	}

	return studentClassroom, nil
}

func (g *GORM) UpdateStudentClassroom(studentClassroom domain.StudentClassroom) error {
	if err := g.DB.Model(&studentClassroom).Where("id = ?", studentClassroom.ID).Update(&studentClassroom).Error; err != nil {
		return err
	}

	return nil
}

func (g *GORM) DeleteStudentClassroomByID(id int64) error {
	var studentClassroom domain.StudentClassroom
	if err := g.DB.Model(&studentClassroom).Where("id = ?", id).Delete(&studentClassroom).Error; err != nil {
		return err
	}

	return nil
}
