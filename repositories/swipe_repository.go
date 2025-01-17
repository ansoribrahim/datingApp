package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SwipeRepository interface {
	RecordSwipe(userID, targetUserID uuid.UUID, isLike bool) error
	GetDailySwipeCount(userID uuid.UUID) (int, error)
	GetSwipedUserIDs(userID uuid.UUID, date time.Time) ([]uuid.UUID, error)
}

type SwipeRepo struct {
	DB *gorm.DB
}

func NewSwipeRepo(db *gorm.DB) *SwipeRepo {
	return &SwipeRepo{DB: db}
}

// RecordSwipe stores the swipe action in the database
func (r *SwipeRepo) RecordSwipe(userID, profileID uuid.UUID, isLike bool) error {
	return r.DB.Exec("INSERT INTO swipes (user_id, profile_id, is_like, created_at) VALUES (?, ?, ?, NOW())",
		userID, profileID, isLike).Error
}

// GetDailySwipeCount returns the number of swipes a user has made today
func (r *SwipeRepo) GetDailySwipeCount(userID uuid.UUID) (int, error) {
	var count int
	err := r.DB.Raw(`
		SELECT COUNT(*) 
		FROM swipes 
		WHERE user_id = ? AND DATE(created_at) = CURRENT_DATE
	`, userID).Scan(&count).Error
	return count, err
}

// GetSwipedUserIDs returns the IDs of users the given user has swiped on today
func (r *SwipeRepo) GetSwipedUserIDs(userID uuid.UUID, date time.Time) ([]uuid.UUID, error) {
	var swipedIDs []uuid.UUID
	err := r.DB.Raw(`
		SELECT target_user_id 
		FROM swipes 
		WHERE user_id = ? AND DATE(created_at) = ?
	`, userID, date.Format("2006-01-02")).Scan(&swipedIDs).Error
	return swipedIDs, err
}
