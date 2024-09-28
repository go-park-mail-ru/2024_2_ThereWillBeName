package models

import "time"

type User struct {
    ID        int64     `json:"id" db:"id"`          				
    Email     string    `json:"email" db:"email"`    
    Password  string    `json:"-" db:"password"`     
    CreatedAt time.Time `json:"created_at" db:"created_at"` 
}
