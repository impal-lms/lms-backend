package services

import (
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func (lms *LMS) GetAllAlumni() ([]domain.Alumni, error) {
	return lms.Repository.GetAllAlumni()
}

func (lms *LMS) CreateAlumni(Alumni domain.Alumni) (domain.Alumni, int, error) {
	Alumni.ID = time.Now().UnixNano()

	Alumni, err := lms.Repository.CreateAlumni(Alumni)
	if err != nil {
		return domain.Alumni{}, 400, errors.Wrap(err, "create Alumni error")
	}

	return Alumni, 201, nil
}

func (lms *LMS) GetAlumniByID(id int64) (domain.Alumni, error) {
	return lms.Repository.GetAlumniByID(id)
}

func (lms *LMS) UpdateAlumni(Alumni domain.Alumni) (domain.Alumni, error) {
	_, err := lms.Repository.GetAlumniByID(Alumni.ID)
	if err != nil {
		return domain.Alumni{}, errors.Wrap(err, "Alumni does not exist")
	}

	return Alumni, lms.Repository.UpdateAlumni(Alumni)
}

func (lms *LMS) DeleteAlumniByID(id int64) (domain.Alumni, error) {
	Alumni, err := lms.Repository.GetAlumniByID(id)
	if err != nil {
		return domain.Alumni{}, errors.Wrap(err, "Alumni does not exist")
	}

	return domain.Alumni{}, lms.Repository.DeleteAlumniByID(Alumni.ID)
}
