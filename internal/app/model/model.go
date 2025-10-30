package model

import (
	"errors"
	"strings"
	"time"
)

type UserID int64

type UserRate struct {
	UserID UserID
	Value  float64
}

type Film struct {
	ID         int64
	Title      string
	Year       int
	Duration   int
	DirectorID int64
	GenreIDs   []int64
	ActorIDs   []int64
	Rates      []UserRate
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (f *Film) Validate() error {
	if strings.TrimSpace(f.Title) == "" {
		return errors.New("Заголовок не может быть пустым")
	}
	if f.Year < 1895 || f.Year > 2025 {
		return errors.New("Год должен быть в промежутке 1895 - 2025")
	}
	if f.Duration <= 0 {
		return errors.New("Продолжительность должна быть больше 0")
	}
	return nil
}

type User struct {
	ID        int64
	Name      string
	Email     string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("Имя не может быть пустым")
	}
	if strings.TrimSpace(u.Email) == "" { //добавить валидацию на адрес
		return errors.New("Email должен быть заполнен")
	}
	return nil
}

type Playlist struct {
	ID        int64
	Title     string
	OwnerID   string
	FilmIDs   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Playlist) Validate() error {
	if strings.TrimSpace(p.Title) == "" {
		return errors.New("Название плейлиста не может быть пустым")
	}
	if p.OwnerID == "" {
		return errors.New("У плейлиста должен быть владелец, нужно добавить ID владельца")
	}
	return nil
}

type Director struct {
	ID        int64
	Firstname string
	Lastname  string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *Director) Validate() error {
	if strings.TrimSpace(d.Firstname) == "" || strings.TrimSpace(d.Lastname) == "" {
		return errors.New("Имя и фамилия режисера должны быть заполнены")
	}
	return nil
}

type Actor struct {
	ID        int64
	Firstname string
	Lastname  string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Actor) Validate() error {
	if strings.TrimSpace(a.Firstname) == "" || strings.TrimSpace(a.Lastname) == "" {
		return errors.New("Имя и фамилия актера должны быть заполнены")
	}
	return nil
}

type Genre struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}

func (g *Genre) Validate() error {
	if strings.TrimSpace(g.Title) == "" {
		return errors.New("Название жанра не может быть пустым")
	}
	return nil
}
