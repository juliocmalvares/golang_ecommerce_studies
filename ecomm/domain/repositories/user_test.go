package repositories

import (
	"context"
	"ecomm/domain/models"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq" // <------------ here
	"github.com/stretchr/testify/assert"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user              = "test_db_user"
	dbName            = "test_db"
	password          = "test_db_password"
	port              = "5432/tcp"
	postgresTimoutSec = 5
	img               = "postgres"
)

var env = map[string]string{
	"POSTGRES_USER":     user,
	"POSTGRES_DB":       dbName,
	"POSTGRES_PASSWORD": password,
}

var db *gorm.DB
var userRepo *UserRepository
var categoryRepo *CategoryRepository
var productRepo *ProductRepository
var productVariationRepo *ProductVariationRepository

func TestMain(m *testing.M) {

	dbUrl := func(host string, port nat.Port) string {
		return getDBUrl(user, password, dbName, port.Port())
	}

	ctx := context.Background()
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        img,
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForSQL(port, "postgres", dbUrl).WithStartupTimeout(time.Second * postgresTimoutSec),
			Name:         "postgres",
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		log.Fatalf("failed to start container: %s", err.Error())
	}

	mappedPort, err := container.MappedPort(ctx, port)
	if err != nil {
		log.Fatalf("failed to get container external port: %s", err.Error())
	}

	log.Printf("postgres container ready and running at port: %s", mappedPort)
	url := getDBUrl(user, password, dbName, mappedPort.Port())
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to establish database connection: %s", err.Error())
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductVariation{})
	userRepo = initUserRepoForTest(db)
	categoryRepo = initCategoryRepoForTest(db)
	productRepo = initProductRepoForTest(db)
	productVariationRepo = initProductVariationRepoForTest(db)

	code := m.Run()

	err = container.Terminate(ctx)
	if err != nil {
		log.Fatalf("failed to terminate container: %s", err.Error())
	}

	os.Exit(code)
}

func getDBUrl(user string, password string, dbName string, port string) string {
	return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, port, dbName)
}

func TestUserRepository_Create(t *testing.T) {
	user := &models.User{
		Email:    "casetest@test.com",
		Password: "123",
	}

	err := db.Create(user).Error

	assert.Nil(t, err)
	assert.Equal(t, user.Email, "casetest@test.com")
	assert.Equal(t, user.Password, "123")
	assert.Equal(t, user.ID, uint(1))

	db.Delete(&user)
}

func TestUserRepository_FindByID(t *testing.T) {
	user := &models.User{
		Email:    "casetest@test.com",
		Password: "123",
	}

	err := db.Create(user).Error
	assert.Nil(t, err)
	userFound, err := userRepo.FindByID(int(user.ID))
	assert.Nil(t, err)
	assert.Equal(t, userFound.Email, "casetest@test.com")
	assert.Equal(t, userFound.Password, "123")
	assert.Equal(t, userFound.ID, uint(2))
	db.Delete(&user)
}

func TestUserRepository_Update(t *testing.T) {
	user := &models.User{
		Email:    "casetest@test.com",
		Password: "123",
	}

	err := db.Create(user).Error
	assert.Nil(t, err)
	user.Password = "456"
	userUpdated, err := userRepo.Update(user)
	assert.Nil(t, err)
	assert.Equal(t, userUpdated.Email, "casetest@test.com")
	assert.Equal(t, userUpdated.Password, "456")
	assert.Equal(t, userUpdated.ID, uint(3))
	db.Delete(&user)
}
