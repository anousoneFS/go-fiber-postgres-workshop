package province_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/anousoneFS/go-fiber-postgres-workshop/province"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var DB *gorm.DB

func TestMain(m *testing.M) {
	dsn := "postgres://oiaaglbm:M7yp7cg1uAG4UpiVazViExpoYwnZTdIw@tiny.db.elephantsql.com/oiaaglbm"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestGetAllProvince(t *testing.T) {
	repo := province.NewRepository(DB)
	p, err := repo.GetAll()
	require.NoError(t, err)
	require.NotEmpty(t, p)
}
