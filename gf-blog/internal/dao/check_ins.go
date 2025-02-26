// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"gf-blog/internal/dao/internal"
)

// internalCheckInsDao is an internal type for wrapping the internal DAO implementation.
type internalCheckInsDao = *internal.CheckInsDao

// checkInsDao is the data access object for the table check_ins.
// You can define custom methods on it to extend its functionality as needed.
type checkInsDao struct {
	internalCheckInsDao
}

var (
	// CheckIns is a globally accessible object for table check_ins operations.
	CheckIns = checkInsDao{
		internal.NewCheckInsDao(),
	}
)

// Add your custom methods and functionality below.
