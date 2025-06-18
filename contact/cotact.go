package contact

import (
	"fmt"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

type Manager struct {
	contacts map[string]Contact
}

// Constructor
func NewManager() *Manager {
	return &Manager{
		contacts: make(map[string]Contact),
	}
}

// Add a new contact
func (m *Manager) AddContact(c Contact) {
	m.contacts[c.Email] = c
}

// Get a contact by email
func (m *Manager) GetContact(email string) (Contact, bool) {
	contact, exists := m.contacts[email]
	return contact, exists
}

// Delete a contact by email
func (m *Manager) DeleteContact(email string) {
	delete(m.contacts, email)
}

// List all contacts
func (m *Manager) ListContacts() {
	if len(m.contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for _, c := range m.contacts {
		fmt.Printf("Name: %s, Phone: %s, Email: %s\n", c.Name, c.Phone, c.Email)
	}
}
