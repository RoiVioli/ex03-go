package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Localisation struct {
	Name string  `json:"name"` // Nom de la commune
	Zip  string  `json:"zip"`  // Code postal de la commune
	Lat  float64 `json:"lat"`  // Latitude de la commune
	Lon  float64 `json:"lon"`  // Longitude de la commune
}

// Fonction principale du programme
func main() {
	maps() // Appelle la fonction maps pour exécuter le code
}

func maps() {
	// Envoie une requête GET à l'API pour récupérer les communes d'un code postal spécifique
	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/zip?zip=E14,GB&appid=c732a4f732342956ec521490b59a7dce")
	if err != nil {
		log.Fatalf("Erreur lors de la récupération de l'API géo : %v", err) // Gérer l'erreur si la requête échoue
	}
	defer resp.Body.Close() // Ferme le corps de la réponse après lecture pour libérer les ressources

	// Vérifie si le code de statut de la réponse est 200 (OK)
	if resp.StatusCode != 200 {
		log.Fatalf("Erreur lors de la récupération de l'API géo, code retour incorrect : %v", resp.StatusCode) // Gérer l'erreur si le code n'est pas 200
	}

	// Lit le corps de la réponse (le contenu des données)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du corps : %v", err) // Gérer l'erreur si la lecture échoue
	}

	// Déclare une variable pour stocker la localisation décodée
	var localisation Localisation // Utiliser un seul objet

	// Décode les données JSON lues dans la variable localisation
	err = json.Unmarshal(body, &localisation)
	if err != nil {
		log.Fatalf("Problème avec le JSON : %v", err) // Gérer l'erreur si le décodage échoue
	}

	// Affiche les informations de la localisation
	log.Printf("Localisation : %v, Code postal : %v, Coordonnées : (%v, %v)", localisation.Name, localisation.Zip, localisation.Lat, localisation.Lon)
}
