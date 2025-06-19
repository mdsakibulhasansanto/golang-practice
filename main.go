package main

import (
	"contact-app/contact"
	reciverfunc "contact-app/contact/reciver_func"
	"fmt"
)

func main() {

	manager := contact.NewManager()

	// Add contacts
	manager.AddContact(contact.Contact{Name: "Santo", Phone: "017xxxxxxx", Email: "santo@example.com"})
	manager.AddContact(contact.Contact{Name: "Anika", Phone: "018xxxxxxx", Email: "anika@example.com"})

	// List all contacts
	fmt.Println("All Contacts:")
	manager.ListContacts()

	// Get a contact
	fmt.Println("\nGetting contact for 'santo@example.com':")
	if c, found := manager.GetContact("santo@example.com"); found {
		fmt.Println("Found:", c)
	} else {
		fmt.Println("Contact not found.")
	}

	// Delete a contact
	fmt.Println("\nDeleting contact 'anika@example.com'")
	manager.DeleteContact("anika@example.com")

	// List again after deletion
	fmt.Println("\nContacts After Deletion:")
	manager.ListContacts()

	per := reciverfunc.Person{
		Name: "Santo",
		Age:  20,
	}

	fmt.Println(per)
	per.Reciver()
	per.Update()

}
