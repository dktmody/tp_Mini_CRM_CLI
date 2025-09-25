package storage

import "errors"

type MemoryStore struct {
	contacts map[int]*Contact
	nextID   int
}

// NewMemoryStore est un constructeur qui initialise proprement notre storer
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *Contact) error {
	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return nil
}

func (ms *MemoryStore) GetAll() ([]*Contact, error) {
	var allContacts []*Contact
	for _, c := range ms.contacts {
		allContacts = append(allContacts, c)
	}
	return allContacts, nil
}

func (ms *MemoryStore) GetByID(id int) (*Contact, error) {
	contact, ok := ms.contacts[id]
	if !ok {
		return nil, errors.New("Contact non trouvé")
	}
	return contact, nil
}

func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	contact, err := ms.GetByID(id)
	if err != nil {
		return err
	}
	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}
	return nil
}

func (ms *MemoryStore) Delete(id int) error {
	if _, ok := ms.contacts[id]; !ok {
		return errors.New("Contact non trouvé")
	}
	delete(ms.contacts, id)
	return nil
}
