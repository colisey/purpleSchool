package link

import (
	"errors"

	"go/adv-demo/pkg/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	// log.Println("CreateRepository")
	// data , err := repo.GetByHash(link.Hash)

	result := repo.Database.DB.Create(link)
	// log.Println("CreateResult")
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link

	// repo.Database.DB.First(&Link, "hash = ? OR id = ?", hash, 10)
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	// fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link

	// repo.Database.DB.First(&Link, "hash = ? OR id = ?", hash, 10)
	result := repo.Database.DB.First(&link, id)
	// fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

// Для удаления в Gorm используется метод Delete. Чтобы ограничить удаляемые записи — применяй Where.

// Пример:
// go
// repo.Database.DB.Where("is_deleted = ?", true).Delete(&User{})

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not found or already deleted")
	}
	return nil
}

// func (repo *LinkRepository) GetByUrl(url string) (*Link, error) {
// 	var link Link

// 	// repo.Database.DB.First(&Link, "hash = ? OR id = ?", hash, 10)
// 	result := repo.Database.DB.Where("url = ?", url).Find(&link)
// 	// result := repo.Database.DB.First(&link, id)
// 	fmt.Println(result)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &link, nil
// }

// Как в GORM применить фильтр для получения записей на основе определенного условия?

// Правильный ответ:

// db.Where("field = ?", value).Find(&objects)

// 📌 Объяснение:
// В GORM правильный порядок вызовов методов:

// db.Where("field = ?", value).Find(&objects)

// Where(...) — добавляет SQL-условие WHERE
// Find(&objects) — выполняет запрос и сохраняет результат в слайс objects
