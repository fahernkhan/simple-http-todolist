package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/postgres" // Driver database PostgreSQL
	"gorm.io/gorm"
)

// User struct dengan JSON tag untuk serialisasi/deserialisasi JSON
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Method untuk User struct
func (u *User) Display() {
	fmt.Printf("User: ID=%d, Name=%s, Email=%s\n", u.ID, u.Name, u.Email)
}

// Interface untuk layanan user
type UserService interface {
	CreateDummyUsers()
	GetAllUsers() []User
}

// Struct implementasi UserService
type UserServiceImpl struct {
	DB *gorm.DB
}

// Implementasi method CreateDummyUsers
func (s *UserServiceImpl) CreateDummyUsers() {
	dummyUsers := []User{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com"},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie Brown", Email: "charlie@example.com"},
	}

	s.DB.Create(&dummyUsers) // Menyimpan dummy data ke database
	fmt.Println("Dummy users created successfully.")
}

// Implementasi method GetAllUsers
func (s *UserServiceImpl) GetAllUsers() []User {
	var users []User
	s.DB.Find(&users) // Mengambil semua data user dari database
	return users
}

func main() {
	// Konfigurasi koneksi PostgreSQL
	dsn := "host=localhost user=postgres password=yourpassword dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{}) // Migrasi tabel

	// Inisialisasi UserService
	userService := &UserServiceImpl{DB: db}

	// Membuat dummy data
	userService.CreateDummyUsers()

	// Mengambil semua user dari database
	users := userService.GetAllUsers()
	fmt.Println("All Users:")
	for _, user := range users {
		user.Display()
	}

	// Serialisasi ke JSON
	userJSON, _ := json.Marshal(users)
	fmt.Println("Users JSON:", string(userJSON))
}
