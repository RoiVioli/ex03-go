package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type communes struct {
	Nom        string `json:"nom"`
	Code       string `json:"code"`
	Population int    `json:"population"`
}

func main() {
	resp, err := http.Get("https://geo.api.gouv.fr/communes?codePostal=12000")
	if err != nil {
		log.Fatalf("Lors de la récupération de geo api : %v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Lors de la récupération de geo api code retour incorrect: %v", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Lecture du body %v", err)
	}
	defer resp.Body.Close() // Ajouté : Ferme le corps de la réponse après lecture

	var communes []communes

	err = json.Unmarshal(body, &communes)
	if err != nil {
		log.Fatalf("Problème avec le json : %v", err) // Ajouté : Affiche l'erreur
	}

	for _, commune := range communes {
		log.Printf("%v avec le code postal %v et la population %v", commune.Nom, commune.Code, commune.Population)
	}
}
