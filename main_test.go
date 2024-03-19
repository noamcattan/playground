package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	db1, err := gorm.Open(sqlite.Open("db1"), &gorm.Config{})
	require.NoError(t, err)
	db2, err := gorm.Open(sqlite.Open("db2"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db1.AutoMigrate(&User{}, &Role{}, &UserRole{}))
	require.NoError(t, db2.AutoMigrate(&UserRole{}, &User{}, &Role{}))
	query := "SELECT sql FROM sqlite_master WHERE tbl_name = 'user_roles'"
	type result struct {
		Sql string
	}
	var res1, res2 result
	db1.Raw(query).Scan(&res1)
	db2.Raw(query).Scan(&res2)
	require.Equal(t, res1.Sql, res2.Sql)
}
