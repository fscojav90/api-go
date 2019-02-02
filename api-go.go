// fscojav90@gmail.com

package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//Modelos
type CharacterDataWrapper struct {
	Code            int                    `json:"code"`
	Status          string                 `json:"status"`
	Copyright       string                 `json:"copyright"`
	AttributionText string                 `json:"atributionText"`
	AttributionHTML string                 `json:"attributionHTML"`
	Data            CharacterDataContainer `json:"data"`
	Etag            string                 `json:"etag"`
}

type CharacterDataContainer struct {
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Count   int         `json:"count"`
	Results []Character `json:"results"`
}

type Character struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Modified    string `json:"modified"`
	Thumbnail   string `json:"thumbnail"`
}

// funciones
//función para obtener el tiempo
func timestamp() string {
	time := time.Now()
	ts := time.Format("20060102150405")
	return ts
}

//función md5
func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//función para buscar por nombre
func search(name string) Character {
	var url = "http://gateway.marvel.com/v1/public/characters?name=" + name + "&ts=" + timestamp() + "&apikey=" + publicKey + "&hash=" + getMD5Hash(timestamp()+privateKey+publicKey)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject CharacterDataWrapper
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Data.Results[0]
}

//función para buscar una lista de personajes (maximo 20)
func searchAll() []Character {
	var url = "http://gateway.marvel.com/v1/public/characters?ts=" + timestamp() + "&apikey=" + publicKey + "&hash=" + getMD5Hash(timestamp()+privateKey+publicKey)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject CharacterDataWrapper
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Data.Results
}

//función para imprimir un personaje
func printCharacter(character Character) {
	fmt.Println("\nid: " + strconv.Itoa(character.Id))
	fmt.Println("name: " + character.Name)
	fmt.Println("description: " + character.Description)
	fmt.Println("modified: " + character.Modified)
}

//función para imprimir una lista de  personajes
func printCharacters(characters []Character) {
	var limite = 0
	if len(characters) >= 20 {
		limite = 20
	} else {
		limite = len(characters)
	}

	for i := 0; i < limite; i++ {
		fmt.Println("\nid: " + strconv.Itoa(characters[i].Id))
		fmt.Println("name: " + characters[i].Name)
		fmt.Println("description: " + characters[i].Description)
		fmt.Println("modified: " + characters[i].Modified)
		fmt.Println("---------------------------------------------------------")
	}
}

// menu principal
func menu() {
	fmt.Println("(1) Buscar por nombre")
	fmt.Println("(2) Listar")

	var opcion string
	fmt.Print("Opción: ")
	fmt.Scanln(&opcion)

	switch opcion {
	case "1":
		var nombre string
		fmt.Print("nombre a buscar: ")
		fmt.Scanln(&nombre)

		printCharacter(search(nombre))
	case "2":
		printCharacters(searchAll())
	default:
		fmt.Println("Opción incorrecta")
	}
}

// constantes
const privateKey = "738460df11a4ee5d6bab89712b9ff6736c81c040"
const publicKey = "c99591d070bd9f94cd2876b9125ffcbb"

func main() {

	menu()
}
