package types

import "time"

// User
type User struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Avatar      string    `json:"avatar"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserPayload struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=3,max=72"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
	UpdateUser(User) error
}

// Client
type Client struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Case
type Case struct {
	ID          int            `json:"id"`
	ClientID    int            `json:"clientId"`
	CaseNumber  string         `json:"caseNumber"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Status      string         `json:"Status"`
	Documents   []CaseDocument `json:"documents"`
	Notes       []CaseNote     `json:"notes"`
	OpenedDate  time.Time      `json:"openedDate"`
	ClosedDate  time.Time      `json:"closedDate"`
	CreatedAt   time.Time      `json:"createdAt"`
}

// Case Document
type CaseDocument struct {
	ID        int       `json:"id"`
	CaseID    int       `json:"caseId"`
	FilePath  string    `json:"filePath"`
	FileName  string    `json:"fileName"`
	FileType  string    `json:"fileType"`
	CreatedAt time.Time `json:"createdAt"`
}

// Case Note
type CaseNote struct {
	ID        int       `json:"id"`
	CaseID    int       `json:"caseId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
