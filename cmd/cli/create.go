package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cette variable stockera l'URL que l'utilisateur va taper
var urlToShorten string

// CreateCmd est la définition de ta commande.
// C'est comme définir un plat sur la carte : son nom, sa description, et ce qui se passe quand on le commande.
var CreateCmd = &cobra.Command{
	Use:   "create", // Le mot clé à taper : ./url-shortener create
	Short: "Crée une URL courte", // Description courte
	Long:  `Crée une URL courte à partir d'une URL longue donnée en argument.`, // Description longue

	// Run est la fonction qui s'exécute quand on tape la commande
	Run: func(cmd *cobra.Command, args []string) {
		// --- ICI commence ton travail de Chef ---

		// 1. On vérifie si l'utilisateur a bien donné une URL
		if urlToShorten == "" {
			fmt.Println("Erreur : Vous devez fournir une URL avec le drapeau --url")
			return // On arrête tout si pas d'URL
		}

		// 2. TEMPORAIRE
		// on affiche juste ce qu'on a reçu pour tester.
		fmt.Printf("Bravo ! J'ai bien reçu l'ordre de raccourcir : %s\n", urlToShorten)
		fmt.Println("TODO: Connecter le LinkService ici plus tard...")

		// Plus tard, on remplaceras les lignes ci-dessus par l'appel au code de tes collègues :
		// newLink, err := linkService.Shorten(urlToShorten)
		// fmt.Printf("URL courte créée : %s\n", newLink.ShortCode)
	},
}

// init est une fonction magique en Go qui s'exécute toute seule au démarrage
func init() {
	// On explique à la commande qu'elle attend un "drapeau" (flag) nommé "url"
	// StringStringVarP permet de lier le drapeau à la variable 'urlToShorten' définie plus haut.
	// Les arguments : (&variable, "nom", "raccourci", "valeur_defaut", "description")
	CreateCmd.Flags().StringVarP(&urlToShorten, "url", "u", "", "L'URL à raccourcir (ex: https://google.com)")
	
	// On dit que ce drapeau est OBLIGATOIRE (optionnel mais recommandé pour la propreté)
	// CreateCmd.MarkFlagRequired("url")
}