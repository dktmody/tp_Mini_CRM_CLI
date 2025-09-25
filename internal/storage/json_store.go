package storage

import (
	"encoding/json"
	"os"
)

type JSONStore struct {
	path     string
	contacts []*Contact
	nextID   int
}

func NewJSONStore(path string) (*JSONStore, error) {
	js := &JSONStore{path: path}
	if err := js.load(); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}
	// compute nextID
	max := 0
	for _, c := range js.contacts {
		if c.ID > max {
			max = c.ID
		}
	}
	js.nextID = max + 1
	return js, nil
}

func (js *JSONStore) load() error {
	f, err := os.Open(js.path)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	return dec.Decode(&js.contacts)
}

func (js *JSONStore) persist() error {
	f, err := os.Create(js.path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(js.contacts)
}

func (js *JSONStore) Add(contact *Contact) error {
	contact.ID = js.nextID
	js.nextID++
	js.contacts = append(js.contacts, contact)
	return js.persist()
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	return js.contacts, nil
}

func (js *JSONStore) GetByID(id int) (*Contact, error) {
	for _, c := range js.contacts {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrContactNotFound(id)
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	c, err := js.GetByID(id)
	if err != nil {
		return err
	}
	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}
	return js.persist()
}

func (js *JSONStore) Delete(id int) error {
	idx := -1
	for i, c := range js.contacts {
		if c.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return ErrContactNotFound(id)
	}
	js.contacts = append(js.contacts[:idx], js.contacts[idx+1:]...)
	return js.persist()
}
