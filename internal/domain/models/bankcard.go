package models

import (
	"errors"
	"time"
	"unicode"
)

type Card struct {
	userID int64
	cardID []rune
	pass   int
	date   time.Time
	meta   string
}

func (c *Card) SetUserID(id int64) {
	c.userID = id
}

func (c *Card) SetCardID(id []rune) error {

	if len(id) != 16 {
		return errors.New("wrong id")
	}
	for i := 0; i < len(id); i++ {
		if !unicode.IsDigit(id[i]) {
			return errors.New("wrong id")
		}
	}

	c.cardID = id

	return nil
}

func (c *Card) SetPass(pass int) error {

	if pass <= 0 || pass >= 1000 {
		return errors.New("wrong card password")
	}

	c.pass = pass

	return nil
}

func (c *Card) SetDate(date time.Time) {
	c.date = date
}

func (c *Card) SetMeta(meta string) {
	c.meta = meta
}

func (c *Card) UserID() int64 {
	return c.userID
}

func (c *Card) CardID() []rune {
	return c.cardID
}

func (c *Card) Pass() int {
	return c.pass
}

func (c *Card) Date() time.Time {
	return c.date
}

func (c *Card) Meta() string {
	return c.meta
}
