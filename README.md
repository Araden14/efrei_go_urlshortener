# 🔗 Go URL Shortener

Bienvenue dans le projet **Go URL Shortener**. Ce projet est une application complète permettant de raccourcir des URLs, de rediriger les utilisateurs instantanément et de suivre les statistiques de clics en temps réel, le tout soutenu par une architecture robuste et asynchrone.

## 📋 Objectifs et Fonctionnalités

Ce service a été conçu pour la performance et la robustesse.

### Fonctionnalités Principales (Core)
* **Raccourcissement d'URLs** : Génération de codes uniques de 6 caractères alphanumériques.
* **Redirection Instantanée** : Redirection HTTP 302 fluide vers l'URL d'origine.
* **Analytics Asynchrones** : Enregistrement des clics via des *Goroutines* et *Channels* (non-bloquant pour l'utilisateur).
* **Surveillance (Monitoring)** : Vérification périodique de la validité des liens (Health Check) avec notifications dans les logs.

### Interfaces
* **API RESTful** : Pour intégrer le service dans d'autres applications.
* **CLI (Ligne de Commande)** : Une interface complète pour gérer le serveur et les liens directement depuis le terminal.

---

## 🛠️ Stack Technique

Ce projet mobilise les technologies suivantes :
* **Langage** : Go (Golang)
* **Interface CLI** : [Cobra](https://github.com/spf13/cobra)
* **API Web** : [Gin Framework](https://github.com/gin-gonic/gin)
* **Base de données** : SQLite avec l'ORM [GORM](https://gorm.io/)
* **Configuration** : [Viper](https://github.com/spf13/viper)

---

## 🚀 Installation et Démarrage

Suivez cette "recette" pour lancer le projet sur votre machine.

### 1. Prérequis
* Go installé sur votre machine (version 1.18+ recommandée).
* Git.

### 2. Installation
Clonez le dépôt et installez les dépendances :

```bash
# Cloner le projet
git clone [https://github.com/Araden14/efrei_go_urlshortener.git](https://github.com/Araden14/efrei_go_urlshortener.git)
cd efrei_go_urlshortener

# Installer les dépendances (nettoyage)
go mod tidy

# Construire l'exécutable
go build -o url-shortener