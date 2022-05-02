package faceit_cc

import (
	"fmt"
	"log"
	"net/mail"
	"time"

	"github.com/jinzhu/gorm"
)

// User is a struct representing the user entity
type User struct {
	Id         string      `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	NickName   string      `gorm:"column:nickname" json:"nickname"`
	Password   string      `json:"password"`
	Email      string      `json:"email"`
	Country    string      `json:"country"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Pagination Pagination  `json:"pagination" gorm:"-"`
	Kafka      Notificater `json:"-" gorm:"-"`
}

type Pagination struct {
	SearchBy    string `json:"search_by"`
	SearchValue string `json:"search_value"`
	ResultsPage int    `json:"search_results_per_page"`
	Offset      int    `json:"offset"`
}

// TableName sets the table name
func (User) TableName() string {
	return "user"
}

// GetById obtains an user using user id property
func (u *User) GetById(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Limit(1).Find(&u).Error; err != nil {
			return err
		}
		return nil
	})
}

// Add adds an user
func (u *User) Add(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Create(u).Error; err != nil {
			return err
		}
		return nil
	})
}

// Update Updates an user
func (u *User) Update(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Model(User{}).Where("id = ?", u.Id).Updates(u).Error; err != nil {
			return err
		}
		return nil
	})
}

// Delete deletes an user
func (u *User) Delete(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Delete(u).Error; err != nil {
			return err
		}
		return nil
	})
}

// List obtains a paginated list of users using Pagination criteria
func (u *User) List(db *gorm.DB, pagination Pagination) ([]User, error) {
	key := fmt.Sprintf("%s = ?", pagination.SearchBy)
	value := pagination.SearchValue
	order := fmt.Sprintf("%s asc", pagination.SearchBy)
	if pagination.SearchBy == "1" {
		order = "id asc"
	}
	offset := pagination.Offset
	if offset > 0 {
		offset = offset * pagination.ResultsPage
	}

	users := []User{}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().
			Limit(pagination.ResultsPage).
			Where(key, value).
			Offset(offset).
			Order(order).
			Find(&users).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return []User{}, err
	}

	log.Println(len(users))

	return users, nil
}

// Validate validates that an user entity has valid property values
func (u *User) Validate() bool {
	if !ValidateUuid(u.Id) {
		return false
	}

	// validate email format
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return false
	}

	// validate other non empty fields (first_name, nickname, password)
	// just to add super simple validation criteria
	if len(u.FirstName) == 0 || len(u.NickName) == 0 || len(u.Password) == 0 {
		return false
	}
	return true
}
