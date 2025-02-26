// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"gf-blog/internal/dao/internal"
)

// internalCategoriesDao is an internal type for wrapping the internal DAO implementation.
type internalCategoriesDao = *internal.CategoriesDao

// categoriesDao is the data access object for the table categories.
// You can define custom methods on it to extend its functionality as needed.
type categoriesDao struct {
	internalCategoriesDao
}

var (
	// Categories is a globally accessible object for table categories operations.
	Categories = categoriesDao{
		internal.NewCategoriesDao(),
	}
)

// Add your custom methods and functionality below.
