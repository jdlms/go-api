package main

var fakeBooks = []*Book{{
	ID:               "1",
	Title:            "7 Habits of Highly Effective People",
	Author:           "Stephen Covey",
	PublishedDate:    "15/08/1989",
	OriginalLanguage: "English",
}}

type fakeStorage struct {
}

func (s fakeStorage) Get(_ string) *Book {
	return fakeBooks[0]
}

func (s fakeStorage) Delete(_ string) *Book {
	return nil
}

func (s fakeStorage) List() []*Book {
	return fakeBooks
}

func (s fakeStorage) Create(_ Book) {
	return
}

func (s fakeStorage) Update(*string, *Book) *Book {
	return fakeBooks[1]
}

// must write test
