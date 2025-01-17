package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email         string    `gorm:"uniqueIndex;not null"`
	PasswordHash  string    `gorm:"not null"`
	Username      string    `gorm:"uniqueIndex;not null"`
	ProfilePicURL string
	IsVerified    bool `gorm:"default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Profile struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"`
	Bio           string
	Interests     string
	LastSwipeDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	User          User           `gorm:"foreignKey:UserID"`
}

type Swipe struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	ProfileID uuid.UUID `gorm:"type:uuid;not null"`
	IsLike    bool      `gorm:"not null"`
	SwipeDate time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User    `gorm:"foreignKey:UserID"`
	Profile   Profile `gorm:"foreignKey:ProfileID"`
}

type PremiumPackage struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PackageName string    `gorm:"not null"`
	Description string
	Price       decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type UserPremium struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	PackageID    uuid.UUID `gorm:"type:uuid;not null"`
	PurchaseDate time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	User         User           `gorm:"foreignKey:UserID"`
	Package      PremiumPackage `gorm:"foreignKey:PackageID"`
}

// BeforeCreate hook to set UUIDs before creation
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (p *Profile) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (s *Swipe) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

func (p *PremiumPackage) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

func (up *UserPremium) BeforeCreate(tx *gorm.DB) error {
	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	return nil
}
