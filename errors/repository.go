package errors

import "errors"

var (
	ErrRepositoryGetAll = errors.New("repository execution of select all query fails")
	ErrRepositoryGetRecord = errors.New("repository execution of select one record query fails")
 	ErrRepositoryCreateRecord = errors.New("repository execution of insert one record query fails")
  ErrRepositoryUpdateRecord = errors.New("repository execution of update one record query fails")
	ErrRepositoryDeleteRecord = errors.New("repository execution of delete one record query fails")
)