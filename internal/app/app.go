package app

import (
	"bufio"
	"fmt"
	"os"
	"projetContact/internal/storage"
	"strconv"
	"strings"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Bienvenue dans Mini CRM v3 !")

	for {
		fmt.Println("\n--- Menu principal ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Modifier un contact")
		fmt.Println("4. Supprimer un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Votre choix : ")

		choice := readUserChoice(reader)

		switch choice {
		case 1:
			handleAddContact(reader, store)
		case 2:
			handleListContacts(store)
		case 3:
			handleUpdateContact(reader, store)
		case 4:
			handleDeleteContact(reader, store)
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer")

		}
	}
}

// Les fonctions "handle..." s'occupent de l'interaction avec l'utilisateur
// et elles appellent la couche de stockage (store) pour effectuer les opérations.
// Elles sont découplées du stockage : elles fonctionnent avec n'importe quel storer

func handleAddContact(reader *bufio.Reader, storer storage.Storer) {
	fmt.Print("Nom du contact : ")
	name := readLine(reader)

	fmt.Print("Email du contact : ")
	email := readLine(reader)

	contact := &storage.Contact{
		Name:  name,
		Email: email,
	}
	err := storer.Add(contact)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Contact '%s' ajouté avec l'ID %d.\n", contact.Name, contact.ID)
}

func handleListContacts(store storage.Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println("Aucun contact à afficher.")
		return
	}

	fmt.Println("\n--- Liste des contacts ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("ID du contact à modifier : ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Printf("Modification de '%s'. Laisser vide pour garder la valeur actuelle.\n", existingContact.Name)

	fmt.Printf("Nouveau nom (%s) : ", existingContact.Name)
	newName := readLine(reader)

	fmt.Printf("Nouvel email (%s) : ", existingContact.Email)
	newEmail := readLine(reader)

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Println("Contact mis à jour avec succès.")
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("ID du contact à supprimer : ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	fmt.Printf("Le contact avec l'ID %d a été supprimé.\n", id)
} // Fonctions utilitaires pour la saisie utilisateur

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readUserChoice(reader *bufio.Reader) int {
	choice, err := strconv.Atoi(readLine(reader))
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	return choice
}

func readInteger(reader *bufio.Reader) int {
	id, err := strconv.Atoi(readLine(reader))
	if err != nil {
		fmt.Println("Erreur : ID invalide. Veuillez entrer un nombre.")
		return -1
	}
	return id
}
