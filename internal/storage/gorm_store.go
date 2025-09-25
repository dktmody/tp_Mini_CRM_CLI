package storage

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type GORMStore struct {
	db *gorm.DB
}

// Close ferme la connexion sous-jacente à la base de données.
func (gs *GORMStore) Close() error {
	sqlDB, err := gs.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func NewGORMStore(path string) (*GORMStore, error) {
	// open a database/sql DB using the modernc sqlite driver (no cgo)
	sqldb, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("Échec ouverture base sqlite : %w", err)
	}

	db, err := gorm.Open(sqlite.New(sqlite.Config{Conn: sqldb}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Échec initialisation GORM : %w", err)
	}

	// Auto migrate Contact schema
	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, fmt.Errorf("Échec auto-migration : %w", err)
	}

	return &GORMStore{db: db}, nil
}

func (gs *GORMStore) Add(contact *Contact) error {
	res := gs.db.Create(contact)
	return res.Error
}

func (gs *GORMStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	res := gs.db.Find(&contacts)
	return contacts, res.Error
}

func (gs *GORMStore) GetByID(id int) (*Contact, error) {
	var c Contact
	res := gs.db.First(&c, id)
	if res.Error != nil {
		return nil, ErrContactNotFound(id)
	}
	return &c, nil
}

func (gs *GORMStore) Update(id int, newName, newEmail string) error {
	c, err := gs.GetByID(id)
	if err != nil {
		return err
	}
	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}
	return gs.db.Save(c).Error
}

func (gs *GORMStore) Delete(id int) error {
	res := gs.db.Delete(&Contact{}, id)
	if res.RowsAffected == 0 {
		return ErrContactNotFound(id)
	}
	return res.Error
}
