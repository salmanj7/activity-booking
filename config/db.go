package config

import (
    "log"
    "github.com/salmanj7/activity-booking/model"
	"golang.org/x/crypto/bcrypt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "youruser:yourpass@tcp(localhost:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to database:", err)
    }

    // AutoMigrate your tables
    db.AutoMigrate(&model.User{})

    // Seed the super admin
    seedSuperAdmin(db)

    DB = db
}

func seedSuperAdmin(db *gorm.DB) {
    var user model.User
    // Check if super admin already exists
    result := db.Where("email = ?", "admin@example.com").First(&user)
    if result.RowsAffected == 0 {
        // Hash the password if needed here
        superAdmin := model.User{
            Name:     "Super Admin",
            Email:    "admin@example.com",
            Password:  HashPassword("supersecurepassword"),
            Role:     "super_admin",
        }
        db.Create(&superAdmin)
        log.Println("✅ Super admin seeded")
    } else {
        log.Println("ℹ️ Super admin already exists")
    }
}

func HashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        log.Fatal(err)
    }
    return string(bytes)
}
