package main

import (
	"fmt"
)

type BookmarkMap map[string]string

func main() {
	// Карта закладок
	bookmarks := make(BookmarkMap, 5)
	// Определяем команды и их описания
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}

}

func getMenu() int {
	var variant int
	commands := BookmarkMap{
		"1": "Посмотреть закладки",
		"2": "Добавить закладку",
		"3": "Удалить закладку",
		"4": "Выход",
	}
	fmt.Println("Введите команду:")
	for key, value := range commands {
		fmt.Println(key, ": ", value)
	}
	fmt.Scan(&variant)
	return variant
}
func printBookmarks(bookmarks BookmarkMap) {
	fmt.Println("Список закладок:")
	for title, url := range bookmarks {
		fmt.Println("%s : %s\n", title, url)
	}
}
func addBookmark(bookmarks BookmarkMap) BookmarkMap {
	var name string
	var url string
	fmt.Print("Введите название закладки:")
	fmt.Scan(&name)
	fmt.Print("Введите адрес закладки:")
	fmt.Scan(&url)
	bookmarks[name] = url
	return bookmarks
}

func deleteBookmark(bookmarks BookmarkMap) BookmarkMap {
	var name string
	fmt.Print("Введите название закладки для удаления:")
	fmt.Scan(&name)
	delete(bookmarks, name)
	return bookmarks
}

// 1 Посмотреть закладки
// 2 Добавить закладку
// 3 Удалить закладку
// 4 Выход
// При 1 - Выводит закладки

// При 2 - 2 поля ввода названия и адреса и после добавление

// При 3 - Ввод названия и удаление по нему

// При 4 - Завершение
