package storage

import (
	"path/filepath"
	"testing"

	"os"
)

func TestGORMStore_CRUD(t *testing.T) {
	td := t.TempDir()
	dbPath := filepath.Join(td, "test_contacts.db")

	// S'assurer que le fichier n'existe pas déjà
	_ = os.Remove(dbPath)

	gs, err := NewGORMStore(dbPath)
	if err != nil {
		t.Fatalf("Échec NewGORMStore: %v", err)
	}
	// fermer la DB à la fin du test pour libérer le fichier
	defer func() {
		_ = gs.Close()
	}()

	// Add
	c := &Contact{Name: "GormTest", Email: "gorm@example.com"}
	if err := gs.Add(c); err != nil {
		t.Fatalf("Échec ajout: %v", err)
	}
	if c.ID == 0 {
		t.Fatalf("ID attendu après ajout, obtenu 0")
	}

	// GetAll
	all, err := gs.GetAll()
	if err != nil {
		t.Fatalf("Échec GetAll: %v", err)
	}
	if len(all) != 1 {
		t.Fatalf("attendu 1 contact, obtenu %d", len(all))
	}

	// GetByID
	got, err := gs.GetByID(c.ID)
	if err != nil {
		t.Fatalf("Échec GetByID: %v", err)
	}
	if got.Name != "GormTest" || got.Email != "gorm@example.com" {
		t.Fatalf("valeurs du contact inattendues : %+v", got)
	}

	// Update
	if err := gs.Update(c.ID, "GormUpdated", "updated@example.com"); err != nil {
		t.Fatalf("Échec Update: %v", err)
	}
	got2, err := gs.GetByID(c.ID)
	if err != nil {
		t.Fatalf("Échec GetByID après update: %v", err)
	}
	if got2.Name != "GormUpdated" || got2.Email != "updated@example.com" {
		t.Fatalf("la mise à jour ne s'est pas propagée : %+v", got2)
	}

	// Delete
	if err := gs.Delete(c.ID); err != nil {
		t.Fatalf("Échec suppression: %v", err)
	}
	// After delete, GetByID should return ErrContactNotFound
	if _, err := gs.GetByID(c.ID); err == nil {
		t.Fatalf("attendu erreur après suppression, obtenu nil")
	}
}
