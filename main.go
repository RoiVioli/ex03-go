package main

import (
	"encoding/json" // Package pour encoder/décoder des données JSON
	"io"            // Package pour les opérations d'entrées/sorties, y compris la lecture de flux
	"log"           // Package pour l'enregistrement des messages de log
	"net/http"      // Package pour effectuer des requêtes HTTP
)

// Définition de la structure communes
type communes struct {
	Nom        string `json:"nom"`        // Nom de la commune
	Code       string `json:"code"`       // Code de la commune
	Population int    `json:"population"` // Population de la commune
}

// Fonction principale du programme
func main() {
	// Envoie une requête GET à l'API pour récupérer les communes d'un code postal spécifique
	resp, err := http.Get("https://geo.api.gouv.fr/communes?codePostal=12000")
	if err != nil {
		log.Fatalf("Lors de la récupération de geo api : %v", err) // Gérer l'erreur si la requête échoue
	}

	// Vérifie si le code de statut de la réponse est 200 (OK)
	if resp.StatusCode != 200 {
		log.Fatalf("Lors de la récupération de geo api code retour incorrect: %v", resp.StatusCode) // Gérer l'erreur si le code n'est pas 200
	}

	// Lit le corps de la réponse (le contenu des données)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Lecture du body %v", err) // Gérer l'erreur si la lecture échoue
	}
	defer resp.Body.Close() // Ferme le corps de la réponse après lecture pour libérer les ressources

	// Déclare une variable pour stocker les communes décodées
	var communes []communes

	// Décode les données JSON lues dans la variable communes
	err = json.Unmarshal(body, &communes)
	if err != nil {
		log.Fatalf("Problème avec le json : %v", err) // Gérer l'erreur si le décodage échoue
	}

	// Boucle sur chaque commune et affiche ses informations
	for _, commune := range communes {
		log.Printf("%v avec le code postal %v et la population %v", commune.Nom, commune.Code, commune.Population)
	}
}
