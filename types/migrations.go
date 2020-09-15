package types

import (
	"errors"
	"fmt"
	"strconv"
)

// NullCreate type ...
type NullCreate struct {
	Name  string
	Valid bool
}

// String func ...
func (create *NullCreate) String() string {
	return fmt.Sprint(*create)
}

// Set func ...
func (create *NullCreate) Set(value string) error {
	if value == "" {
		return errors.New("cannot be empty")
	}

	create.Name = value
	create.Valid = true

	return nil
}

// NullDownTo type ...
type NullDownTo struct {
	Version int64
	Valid   bool
}

// String func ...
func (downto *NullDownTo) String() string {
	return fmt.Sprint(*downto)
}

// Set func ...
func (downto *NullDownTo) Set(value string) error {
	i, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return err
	}

	downto.Version = i
	downto.Valid = true

	return nil
}

var MigrateValues = SuperSlice{
	"up", "up-by-one", "down",
	"redo", "reset",
	"status", "version",
}

// NullMigrate type ...
type NullMigrate struct {
	Migrate string
	Valid   bool
}

// String func ...
func (migrate *NullMigrate) String() string {
	return fmt.Sprint(*migrate)
}

// Set func ...
func (migrate *NullMigrate) Set(value string) error {
	if _, found := MigrateValues.Find(value); found {
		migrate.Migrate = value
		migrate.Valid = true

		return nil
	}

	return errors.New("not found in all possible values")
}

// NullUpTo type ...
type NullUpTo struct {
	Version int64
	Valid   bool
}

// String func ...
func (upto *NullUpTo) String() string {
	return fmt.Sprint(*upto)
}

// Set func ...
func (upto *NullUpTo) Set(value string) error {
	i, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		return err
	}

	upto.Version = i
	upto.Valid = true

	return nil
}
