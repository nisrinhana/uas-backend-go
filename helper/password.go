package helper

import "golang.org/x/crypto/bcrypt"

// HashPassword menerima password plaintext dan mengembalikan hashed string
func HashPassword(password string) string {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err) // atau return error sesuai kebutuhan
    }
    return string(hash)
}

// ComparePassword membandingkan password plaintext dengan hashed password
func ComparePassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
