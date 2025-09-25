package storage

import (
	"os"
	"testing"
)

func TestMemoryStoreCRUD(t *testing.T) {
	ms := NewMemoryStore()

	c := &Contact{Name: "T1", Email: "t1@example.com"}
	if err := ms.Add(c); err != nil {
		t.Fatalf("Échec ajout: %v", err)
	}
	if c.ID == 0 {
		t.Fatalf("ID attendu non nul")
	}

	all, err := ms.GetAll()
	if err != nil {
		t.Fatalf("Échec GetAll: %v", err)
	}
	if len(all) != 1 {
		t.Fatalf("attendu 1 contact, obtenu %d", len(all))
	}

	got, err := ms.GetByID(c.ID)
	if err != nil {
		t.Fatalf("Échec GetByID: %v", err)
	}
	if got.Name != c.Name {
		t.Fatalf("attendu nom %s, obtenu %s", c.Name, got.Name)
	}

	if err := ms.Update(c.ID, "T1b", "t1b@example.com"); err != nil {
		t.Fatalf("Échec Update: %v", err)
	}
	got2, _ := ms.GetByID(c.ID)
	if got2.Email != "t1b@example.com" {
		t.Fatalf("update didn't change email")
	}

	if err := ms.Delete(c.ID); err != nil {
		t.Fatalf("Échec Delete: %v", err)
	}
}

func TestJSONStoreCRUD(t *testing.T) {
	path := "test_store.json"
	// ensure cleanup
	_ = os.Remove(path)
	js, err := NewJSONStore(path)
	if err != nil {
		t.Fatalf("Échec NewJSONStore: %v", err)
	}

	c := &Contact{Name: "J1", Email: "j1@example.com"}
	if err := js.Add(c); err != nil {
		t.Fatalf("Échec ajout: %v", err)
	}
	if c.ID == 0 {
		t.Fatalf("ID attendu non nul")
	}

	all, err := js.GetAll()
	if err != nil {
		t.Fatalf("Échec GetAll: %v", err)
	}
	if len(all) != 1 {
		t.Fatalf("attendu 1 contact, obtenu %d", len(all))
	}

	got, err := js.GetByID(c.ID)
	if err != nil {
		t.Fatalf("Échec GetByID: %v", err)
	}
	if got.Name != c.Name {
		t.Fatalf("attendu nom %s, obtenu %s", c.Name, got.Name)
	}

	if err := js.Update(c.ID, "J1b", "j1b@example.com"); err != nil {
		t.Fatalf("Échec Update: %v", err)
	}
	got2, _ := js.GetByID(c.ID)
	if got2.Email != "j1b@example.com" {
		t.Fatalf("update didn't change email")
	}

	if err := js.Delete(c.ID); err != nil {
		t.Fatalf("Échec Delete: %v", err)
	}
	_ = os.Remove(path)
}
