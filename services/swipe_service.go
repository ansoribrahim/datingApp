package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"datingApp/models"
	"datingApp/repositories"
)

type SwipeService struct {
	UserRepo  repositories.UserRepository
	SwipeRepo repositories.SwipeRepository
}

func NewSwipeService(userRepo repositories.UserRepository, swipeRepo repositories.SwipeRepository) *SwipeService {
	return &SwipeService{
		UserRepo:  userRepo,
		SwipeRepo: swipeRepo,
	}
}

// SwipeRight handles a "like" action
func (s *SwipeService) SwipeRight(userID, profileID uuid.UUID) error {
	if userID == profileID {
		return errors.New("cannot swipe on your own profile")
	}

	// Check daily swipe quota
	count, err := s.SwipeRepo.GetDailySwipeCount(userID)
	if err != nil {
		return err
	}
	if count >= 10 {
		return errors.New("daily swipe quota exceeded")
	}

	// Record swipe
	return s.SwipeRepo.RecordSwipe(userID, profileID, true)
}

// SwipeLeft handles a "pass" action
func (s *SwipeService) SwipeLeft(userID, profileID uuid.UUID) error {
	if userID == profileID {
		return errors.New("cannot swipe on your own profile")
	}

	// Check daily swipe quota
	count, err := s.SwipeRepo.GetDailySwipeCount(userID)
	if err != nil {
		return err
	}
	if count >= 10 {
		return errors.New("daily swipe quota exceeded")
	}

	// Record swipe
	return s.SwipeRepo.RecordSwipe(userID, profileID, false)
}

// GetPotentialMatches returns profiles that the user has not swiped on yet today
func (s *SwipeService) GetPotentialMatches(userID uuid.UUID) ([]models.User, error) {
	swipedIDs, err := s.SwipeRepo.GetSwipedUserIDs(userID, time.Now())
	if err != nil {
		return nil, err
	}

	return s.UserRepo.GetUnswipedUsers(userID, swipedIDs)
}
