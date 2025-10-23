package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kevinmso/estudos-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool

	GetCustomersByEmail(ctx context.Context, email string) ([]models.Customer, error)
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
	GetProductsByVendor(ctx context.Context, vendorId string) ([]models.Product, error)
	GetAllServices(ctx context.Context) ([]models.Service, error)

	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	AddService(ctx context.Context, service *models.Service) (*models.Service, error)
	AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {
	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Não foi possível carregar .env: %v", err)
	}

	// Ler variáveis
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Montar DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	// Conectar ao banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	client := Client{
		DB: db,
	}

	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 AS ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	return ready == "1"
}
