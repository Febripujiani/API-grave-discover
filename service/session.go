package service

import (
	"github.com/Febripujiani/model"
	repo "github.com/Febripujiani/repository"
)

type SessionService interface {
	GetSessionByUsername(username string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByUsername(username string) (model.Session, error) {
	session, err := c.sessionRepo.SessionAvailUsername(username)
	if err != nil {
		return model.Session{}, err
	}
	return session, nil
	 // TODO: replace this
}

func (s *sessionService) GetSessionsByRole(role string) ([]model.Session, error) {
	sessions, err := s.sessionRepo.GetSessionsByRole(role)
	if err != nil {
	return []model.Session{}, err
	}
	return sessions, nil
	}