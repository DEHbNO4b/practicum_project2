package models

import (
	"errors"
	"unicode"
)

type Card struct {
	userID int64
	cardID []rune
	pass   string
	date   string
	meta   string
}

func (c *Card) SetUserID(id int64) {
	c.userID = id
}

func (c *Card) SetCardID(id []rune) error {

	// if len(id) != 16 {
	// 	return errors.New("wrong id")
	// }
	// for i := 0; i < len(id); i++ {
	// 	if !unicode.IsDigit(id[i]) {
	// 		return errors.New("wrong id")
	// 	}
	// }

	c.cardID = id

	return nil
}

func (c *Card) SetPass(pass string) error {

	if len(pass) != 3 {
		return errors.New("wrong card password len")
	}
	for i := 0; i < len(pass); i++ {
		if !unicode.IsDigit(rune(pass[i])) {
			return errors.New("wrong id")
		}
	}
	c.pass = pass

	return nil
}

func (c *Card) SetDate(date string) {
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

func (c *Card) Pass() string {
	return c.pass
}

func (c *Card) Date() string {
	return c.date
}

func (c *Card) Meta() string {
	return c.meta
}

func (c *Card) String() string {

	str := "card id: " + string(c.cardID) + "\n"
	str += "password: " + c.pass + "\n"
	str += "date: " + c.date + "\n"
	str += "info: " + c.meta + "\n"

	return str
}
