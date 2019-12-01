package services

import (
	"time"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/pkg/errors"
)

func (lms *LMS) GetAllPresence(studentID int64) ([]domain.Presence, error) {
	return lms.Repository.GetAllPresence(studentID)
}

func (lms *LMS) CreatePresence(Presence domain.Presence) (domain.Presence, int, error) {
	Presence.ID = time.Now().UnixNano()

	Presence, err := lms.Repository.CreatePresence(Presence)
	if err != nil {
		return domain.Presence{}, 400, errors.Wrap(err, "create Presence error")
	}

	return Presence, 201, nil
}

func (lms *LMS) GetPresenceByID(id int64) (domain.Presence, error) {
	return lms.Repository.GetPresenceByID(id)
}

func (lms *LMS) UpdatePresence(Presence domain.Presence) (domain.Presence, error) {
	_, err := lms.Repository.GetPresenceByID(Presence.ID)
	if err != nil {
		return domain.Presence{}, errors.Wrap(err, "Presence does not exist")
	}

	return Presence, lms.Repository.UpdatePresence(Presence)
}

func (lms *LMS) DeletePresenceByID(id int64) (domain.Presence, error) {
	Presence, err := lms.Repository.GetPresenceByID(id)
	if err != nil {
		return domain.Presence{}, errors.Wrap(err, "Presence does not exist")
	}

	return domain.Presence{}, lms.Repository.DeletePresenceByID(Presence.ID)
}
