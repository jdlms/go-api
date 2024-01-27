package main

import "go-api/pkg/models"

var fakeBooks = []*models.Book{{
	ID:               "1",
	Title:            "7 Habits of Highly Effective People",
	Author:           "Stephen Covey",
	PublishedDate:    "15/08/1989",
	OriginalLanguage: "English",
}}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) *models.Book {
	return fakeBooks[0]
}

func (s fakeStorage) Delete(_ string) *models.Book {
	return nil
}

func (s fakeStorage) List() []*models.Book {
	return fakeBooks
}

// func (s fakeStorage) Create(models.Book {
// 	return
// }

func (s fakeStorage) Update(*string, models.Book) *models.Book {
	return fakeBooks[1]
}

// must write test
