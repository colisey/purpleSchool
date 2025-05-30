package product

// package user

type Product struct {
	ID    uint
	Title string
	Price float64
}

func (Product) TableName() string {
	return "shop_products"
}

// product := Product{Name: "Notebook", Price: 999.99}
// result := db.Create(&product)

// if result.Error != nil {
//     log.Println("Ошибка при сохранении:", result.Error)
// }

// var product Product
// result := db.First(&product, 1)

// var products []Product
// result := db.Where("price > ?", 500).Find(&products)

// Теперь Gorm всегда будет использовать таблицу accounts
// для этой модели, даже если структура называется User.

// Давай закрепим — напиши объявление структуры Product с полями
// ID (uint), Title (string) и цену Price (float64),
// у которой таблица будет называться 'shop_products'.

//	type Author struct {
//		ID    uint
//		Name  string
//		Books []Book // связь has many
//	}
// type Category struct {
// 	ID       uint
// 	Title    string
// 	Products []Product // один ко многим
// }

// type Product struct {
// 	ID         uint
// 	Title      string
// 	Price      float64
// 	CategoryID uint // внешний ключ
// }

// err := db.AutoMigrate(&Category{}, &Product{})
// if err != nil {
// 	log.Fatalf("Ошибка миграции: %v", err)
// }

// TODO как в функции миграции (AutoMigrate) правильно выполнить миграцию
// с такими связанными моделями, чтобы создались обе таблицы и внешний ключ?
// Приведи пример вызова.

// GORM создаст таблицу categories, затем products.
// Обнаружит поле CategoryID uint в Product — автоматически добавит внешний ключ (FOREIGN KEY) к categories(id).
// Также благодаря Products []Product в Category GORM поймёт, что это связь "один ко многим" (One-To-Many).

// Это стандарт для Gorm — теперь один Author будет иметь срез Books, связанных через AuthorID.

// Попробуй теперь самостоятельно объявить структуру Category
// и структуру Product так, чтобы у одной категории было много продуктов (has many).
