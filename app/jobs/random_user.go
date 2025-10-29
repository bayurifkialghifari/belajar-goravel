package jobs

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/goravel/framework/facades"
	"karuhundeveloper.com/gogo/app/models"
)

type RandomUser struct {
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Signature The name and signature of the job.
func (receiver *RandomUser) Signature() string {
	return "random_user"
}

// Generate a random string of given length
func generateRandomString(length int) string {
	// seededRand is a local random number generator instance
	var seede *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seede.Intn(len(charset))]
	}

	// Unix timestamp to ensure uniqueness
	timespamp := time.Now().Unix()
	append := strconv.Itoa(int(timespamp))

	return string(b) + append
}

// Handle Execute the job.
func (receiver *RandomUser) Handle(args ...any) error {
	// Create a random email
	password, _ := facades.Hash().Make("password")
	userData := &models.User{
		Name:     args[0].(string),
		Email:    generateRandomString(5) + "@example.com",
		Password: password,
	}

	// Save the user to the database
	if err := facades.Orm().Query().Model(&models.User{}).Create(userData); err != nil {
		return err
	}

	// Sleep for 30 seconds to simulate a long-running task
	// time.Sleep(30 * time.Second)

	return nil
}

// ShouldRetry determines if the job should be retried based on the error.
func (r *RandomUser) ShouldRetry(err error, attempt int) (retryable bool, delay time.Duration) {
  return true, 10 * time.Second
}