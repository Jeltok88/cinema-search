package model

import (
	"testing"
	"time"
)

func TestFilm(t *testing.T) {
	tests := []struct {
		name     string
		film     Film
		expected Film
	}{
		{
			name: "complete film data",
			film: Film{
				ID:         1,
				Title:      "Интерстеллар",
				Year:       2014,
				Duration:   169,
				DirectorID: 100,
				GenreIDs:   []int64{1, 2},
				ActorIDs:   []int64{10, 20, 30},
				Rates: []UserRate{
					{UserID: 1, Value: 4.5},
					{UserID: 2, Value: 5.0},
				},
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			expected: Film{
				ID:         1,
				Title:      "Интерстеллар",
				Year:       2014,
				Duration:   169,
				DirectorID: 100,
				GenreIDs:   []int64{1, 2},
				ActorIDs:   []int64{10, 20, 30},
				Rates: []UserRate{
					{UserID: 1, Value: 4.5},
					{UserID: 2, Value: 5.0},
				},
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "minimal film data",
			film: Film{
				ID:         2,
				Title:      "Тестовый фильм",
				Year:       2023,
				Duration:   120,
				DirectorID: 200,
				GenreIDs:   []int64{},
				ActorIDs:   []int64{},
				Rates:      []UserRate{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			},
			expected: Film{
				ID:         2,
				Title:      "Тестовый фильм",
				Year:       2023,
				Duration:   120,
				DirectorID: 200,
				GenreIDs:   []int64{},
				ActorIDs:   []int64{},
				Rates:      []UserRate{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			},
		},
		{
			name: "film with multiple genres and actors",
			film: Film{
				ID:         3,
				Title:      "Большой фильм",
				Year:       2020,
				Duration:   180,
				DirectorID: 300,
				GenreIDs:   []int64{1, 2, 3, 4, 5},
				ActorIDs:   []int64{10, 11, 12, 13, 14, 15},
				Rates: []UserRate{
					{UserID: 1, Value: 4.8},
					{UserID: 2, Value: 3.5},
					{UserID: 3, Value: 4.2},
				},
				CreatedAt: time.Date(2020, 5, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 5, 11, 0, 0, 0, 0, time.UTC),
			},
			expected: Film{
				ID:         3,
				Title:      "Большой фильм",
				Year:       2020,
				Duration:   180,
				DirectorID: 300,
				GenreIDs:   []int64{1, 2, 3, 4, 5},
				ActorIDs:   []int64{10, 11, 12, 13, 14, 15},
				Rates: []UserRate{
					{UserID: 1, Value: 4.8},
					{UserID: 2, Value: 3.5},
					{UserID: 3, Value: 4.2},
				},
				CreatedAt: time.Date(2020, 5, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2020, 5, 11, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "film with zero values",
			film: Film{
				ID:         0,
				Title:      "",
				Year:       0,
				Duration:   0,
				DirectorID: 0,
				GenreIDs:   nil,
				ActorIDs:   nil,
				Rates:      nil,
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			},
			expected: Film{
				ID:         0,
				Title:      "",
				Year:       0,
				Duration:   0,
				DirectorID: 0,
				GenreIDs:   nil,
				ActorIDs:   nil,
				Rates:      nil,
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			},
		},
		{
			name: "film with decimal ratings",
			film: Film{
				ID:         4,
				Title:      "Фильм с дробными рейтингами",
				Year:       2022,
				Duration:   150,
				DirectorID: 400,
				GenreIDs:   []int64{1},
				ActorIDs:   []int64{40},
				Rates: []UserRate{
					{UserID: 1, Value: 4.7},
					{UserID: 2, Value: 3.2},
					{UserID: 3, Value: 4.9},
				},
				CreatedAt: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 16, 0, 0, 0, 0, time.UTC),
			},
			expected: Film{
				ID:         4,
				Title:      "Фильм с дробными рейтингами",
				Year:       2022,
				Duration:   150,
				DirectorID: 400,
				GenreIDs:   []int64{1},
				ActorIDs:   []int64{40},
				Rates: []UserRate{
					{UserID: 1, Value: 4.7},
					{UserID: 2, Value: 3.2},
					{UserID: 3, Value: 4.9},
				},
				CreatedAt: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 3, 16, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Проверяем основные поля
			if tt.film.ID != tt.expected.ID {
				t.Errorf("ID = %d, want %d", tt.film.ID, tt.expected.ID)
			}
			if tt.film.Title != tt.expected.Title {
				t.Errorf("Title = %s, want %s", tt.film.Title, tt.expected.Title)
			}
			if tt.film.Year != tt.expected.Year {
				t.Errorf("Year = %d, want %d", tt.film.Year, tt.expected.Year)
			}
			if tt.film.Duration != tt.expected.Duration {
				t.Errorf("Duration = %d, want %d", tt.film.Duration, tt.expected.Duration)
			}
			if tt.film.DirectorID != tt.expected.DirectorID {
				t.Errorf("DirectorID = %d, want %d", tt.film.DirectorID, tt.expected.DirectorID)
			}

			// Проверяем слайсы
			if len(tt.film.GenreIDs) != len(tt.expected.GenreIDs) {
				t.Errorf("GenreIDs length = %d, want %d", len(tt.film.GenreIDs), len(tt.expected.GenreIDs))
			} else {
				for i, genreID := range tt.film.GenreIDs {
					if genreID != tt.expected.GenreIDs[i] {
						t.Errorf("GenreIDs[%d] = %d, want %d", i, genreID, tt.expected.GenreIDs[i])
					}
				}
			}

			if len(tt.film.ActorIDs) != len(tt.expected.ActorIDs) {
				t.Errorf("ActorIDs length = %d, want %d", len(tt.film.ActorIDs), len(tt.expected.ActorIDs))
			} else {
				for i, actorID := range tt.film.ActorIDs {
					if actorID != tt.expected.ActorIDs[i] {
						t.Errorf("ActorIDs[%d] = %d, want %d", i, actorID, tt.expected.ActorIDs[i])
					}
				}
			}

			// Проверяем рейтинги
			if len(tt.film.Rates) != len(tt.expected.Rates) {
				t.Errorf("Rates length = %d, want %d", len(tt.film.Rates), len(tt.expected.Rates))
			} else {
				for i, rate := range tt.film.Rates {
					if rate.UserID != tt.expected.Rates[i].UserID {
						t.Errorf("Rates[%d].UserID = %d, want %d", i, rate.UserID, tt.expected.Rates[i].UserID)
					}
					if rate.Value != tt.expected.Rates[i].Value {
						t.Errorf("Rates[%d].Value = %.2f, want %.2f", i, rate.Value, tt.expected.Rates[i].Value)
					}
				}
			}

			// Проверяем временные метки
			if !tt.film.CreatedAt.Equal(tt.expected.CreatedAt) {
				t.Errorf("CreatedAt = %v, want %v", tt.film.CreatedAt, tt.expected.CreatedAt)
			}
			if !tt.film.UpdatedAt.Equal(tt.expected.UpdatedAt) {
				t.Errorf("UpdatedAt = %v, want %v", tt.film.UpdatedAt, tt.expected.UpdatedAt)
			}
		})
	}
}
