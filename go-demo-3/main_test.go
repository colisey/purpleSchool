package main

import (
	"testing"
)

// TestBookmarksAdd проверяет добавление закладок
func TestBookmarksAdd(t *testing.T) {
	bookmarks := NewBookmarks()

	// Тест 1: Добавление первой закладки
	bookmarks.Add("Google", "https://google.com")
	if len(bookmarks.items) != 1 {
		t.Errorf("Ожидалось 1 закладка, получено %d", len(bookmarks.items))
	}
	if bookmarks.items[0].Name != "Google" {
		t.Errorf("Ожидалось название 'Google', получено '%s'", bookmarks.items[0].Name)
	}
	if bookmarks.items[0].URL != "https://google.com" {
		t.Errorf("Ожидался URL 'https://google.com', получен '%s'", bookmarks.items[0].URL)
	}

	// Тест 2: Добавление второй закладки
	bookmarks.Add("GitHub", "https://github.com")
	if len(bookmarks.items) != 2 {
		t.Errorf("Ожидалось 2 закладки, получено %d", len(bookmarks.items))
	}
}

// TestBookmarksRemove проверяет удаление закладок
func TestBookmarksRemove(t *testing.T) {
	bookmarks := NewBookmarks()

	// Подготовка данных
	bookmarks.Add("Google", "https://google.com")
	bookmarks.Add("GitHub", "https://github.com")

	// Тест 1: Удаление существующей закладки
	bookmarks.Remove("Google")
	if len(bookmarks.items) != 1 {
		t.Errorf("Ожидалось 1 закладка, получено %d", len(bookmarks.items))
	}
	if bookmarks.items[0].Name != "GitHub" {
		t.Errorf("Ожидалось название 'GitHub', получено '%s'", bookmarks.items[0].Name)
	}

	// Тест 2: Удаление несуществующей закладки
	bookmarks.Remove("NonExistent")
	if len(bookmarks.items) != 1 {
		t.Errorf("Количество закладок не должно измениться, получено %d", len(bookmarks.items))
	}
}

// TestBookmarksList проверяет вывод списка закладок
func TestBookmarksList(t *testing.T) {
	bookmarks := NewBookmarks()

	// Тест 1: Пустой список
	if len(bookmarks.items) != 0 {
		t.Error("Список закладок должен быть пустым")
	}

	// Тест 2: Список с закладками
	bookmarks.Add("Google", "https://google.com")
	bookmarks.Add("GitHub", "https://github.com")

	if len(bookmarks.items) != 2 {
		t.Errorf("Ожидалось 2 закладки, получено %d", len(bookmarks.items))
	}

	// Проверка содержимого
	if bookmarks.items[0].Name != "Google" || bookmarks.items[0].URL != "https://google.com" {
		t.Error("Первая закладка содержит неверные данные")
	}
	if bookmarks.items[1].Name != "GitHub" || bookmarks.items[1].URL != "https://github.com" {
		t.Error("Вторая закладка содержит неверные данные")
	}
}
