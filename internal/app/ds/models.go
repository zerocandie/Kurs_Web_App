package ds

import "time"

type UserRole string

const (
	RoleCitizen   UserRole = "citizen"
	RoleModerator UserRole = "moderator"
)

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	Login string `gorm:"type:varchar(100);unique;not null" json:"login"`

	Password string `gorm:"type:varchar(255);not null" json:"-"`

	FullName string `gorm:"type:varchar(255);not null" json:"full_name"`

	Email string `gorm:"type:varchar(255);unique" json:"email"`

	Role UserRole `gorm:"type:varchar(20);default:'citizen'" json:"role"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

/* ===========================
   CITY
=========================== */

type City struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	Title string `gorm:"type:varchar(150);not null" json:"title"`

	Region string `gorm:"type:varchar(150);not null" json:"region"`

	Events []Event `gorm:"foreignKey:CityID" json:"events,omitempty"`
}

func (City) TableName() string {
	return "cities"
}

/* ===========================
   EVENT
=========================== */

type Event struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	Title string `gorm:"type:varchar(255);not null" json:"title"`

	Description string `gorm:"type:text" json:"description"`

	Image string `gorm:"type:varchar(500)" json:"image"`

	Video string `gorm:"type:varchar(500)" json:"video"`

	EventDate time.Time `gorm:"not null;index" json:"event_date"`

	Location string `gorm:"type:varchar(255)" json:"location"`

	CityID uint `gorm:"not null;index" json:"city_id"`

	City City `gorm:"foreignKey:CityID" json:"city"`

	IsActive bool `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Event) TableName() string {
	return "events"
}

/* ===========================
   VOTE
=========================== */

type VoteStatus string

const (
	VoteDraft     VoteStatus = "draft"
	VoteFormed    VoteStatus = "formed"
	VoteCompleted VoteStatus = "completed"
	VoteRejected  VoteStatus = "rejected"
	VoteDeleted   VoteStatus = "deleted"
)

type Vote struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	Status VoteStatus `gorm:"type:varchar(30);default:'draft';index" json:"status"`

	CreatorID uint `gorm:"not null;index" json:"creator_id"`

	Creator User `gorm:"foreignKey:CreatorID" json:"creator"`

	ModeratorID *uint `gorm:"index" json:"moderator_id"`

	Moderator *User `gorm:"foreignKey:ModeratorID" json:"moderator,omitempty"`

	Region string `gorm:"type:varchar(255)" json:"region"`

	Result int `gorm:"default:0" json:"result"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	FormedAt *time.Time `json:"formed_at,omitempty"`

	CompletedAt *time.Time `json:"completed_at,omitempty"`

	VoteItems []VoteItem `gorm:"foreignKey:VoteID" json:"vote_items,omitempty"`
}

func (Vote) TableName() string {
	return "votes"
}

/* ===========================
   MANY TO MANY
=========================== */

type VoteItem struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	VoteID uint `gorm:"not null;index" json:"vote_id"`

	Vote Vote `gorm:"foreignKey:VoteID" json:"-"`

	EventID uint `gorm:"not null;index" json:"event_id"`

	Event Event `gorm:"foreignKey:EventID" json:"event"`

	Priority int `gorm:"default:1" json:"priority"`

	Value int `gorm:"default:1" json:"value"`

	Result string `gorm:"type:text" json:"result"`
}

func (VoteItem) TableName() string {
	return "vote_items"
}
