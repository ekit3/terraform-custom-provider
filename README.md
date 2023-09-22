# Terraform, provider personalisé ou comment consommer de l'api autrement

Terraform est devenu un outil imanquable dans le monde du devops, et plus particulièrement dans tout ce qui est infrastructure as code voir uniquement dans celui ci.  
Dans cet article, nous allons mettre en evidence une autre manière de penser l'utilisation de terraform, le "service as code". On oublie souvent que terraform fonctionnent grace aux différentes apis fournient par des tiers. Mais comment tout cela fonctionne ? C'est ce que nous allons voir ici, à la fin de cette article, vous aurez développez un provider terraform simple, connecté à une api "maison"

*Pour mieux appréhender cet article, il est préférable d’être déjà initié à Terraform et de manière plus générale une connaissance de l’infrastructure as code. Ne vous inquiétez pas, même si vous n’êtes pas familier avec ces concepts, nous allons tout de même commencer par des petits rappels, qui vous permettront de suivre de guide pas à pas.*

## Petit rappel

### Terraform

Terraform est un outil qui va permettre de créer, mettre à jour & versionner vos infrastructures de manière sécurisée et efficiente. Une infrastructure composée de ressources cloud, comme des bases de données, des machines virtuelles ou n’importe quel service de votre fournisseur cloud préféré.  
Il peut aussi déployer des ressources “maison”, via du développement custom, qui seront propre à vos besoins et c'est que nous allons voir ici ! :)  

### Terraform provider

Les providers terraform vont vous permettre d'intéragir avec un service tiers, type GCP ou Github. On pourrait comparer les providers terraform avec des librairies dans les langages de développement. Ces intéractions seront réaliser via les différentes ressources fournis par ces providers, par exemple le provider Terraform de Github fournis une ressource "github_team" qui permet de créer, modifier ou supprimer une team github.  

*Si vous n'avais jamais manipuler de terraform, je vous conseil d'aller voir ce [guide](https://github.com/ekit3/terraform-scaleway) dans un premier temps avant de continuer celui ci*

## Mais a quoi bon un provider custom ?

Comme je l'ai dis en introduction, on peut voir terraform comme une autre manière de consommer de l'API. Dans la plupart des cas, quand on veut consommer une api, que ca soit pour de la lecture ou de l'écriture, on va développer soit une autre API, soit du batch ce qui demande déjà ( selon les besoins) une certaine expertise et un certain nombre de resources( temps, materielles, humaines), cela peut vite devenir un frein a la consommation de vos services.  
Avec un provider propre à vos services, vous offrez la possibilité à vos utilisateur de consommer directement et simplement votre ou vos Apis, le code terraform étant très descriptif.  
En plus de cela vous bénificiez des autres avantages de terraform, l'état de votre utilisation des services correspond a votre code, la persistance de l'état des objets crées plustot pratique dans les cas de DRP par exemple.

## J’en veux un, comment je fais ?

Ici nous allons réaliser un guide pas à pas, pour développer ce fameux provider custom, pour cela dans le repository git, il y a un fichier docker-compose.yml qui va vous permettre de lancer une bdd, une api et un petit front pour visualiser le contenue de la bdd.
Cette api n'est qu'un simple crud sur un objet "Cour", qui possède plusieurs champs : un nom, une durée & une description

### Prérequis

- [Terraform CLI](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
- [Docker Compose](https://docs.docker.com/compose/install/#scenario-one-install-docker-desktop)
- [Golang](https://go.dev/doc/install)

### On code !!

Hashicorp fournit différente librairies pour faciliter le développement de provider terraform, il fournit egalement des repositories d exemples qui permettent de démarrer rapidement le développement d'un provider et tout en ensemble de guide très complet.

#### Pas à pas

cloner le dossier de démarrage

expliquer l'arborescence du repo

faire démarrer le compose, faire un test d'accès a l'app

Leur faire démarrer sur le dossier resources/terraform-provider-courses ( fichiers de démarrage ? jsais pas traduire boilerplate)

expliquer le contenu initial du main.go & du provider.go,les lib hashicorp et la lib 'custom'
expliquer le main.tf ( surtout l'import, qui correspond au go mod init "nom dossier"/sous dossier
)

le detail du makefile ( le lien entre le nom qu'on donne au binaire, comment on l'importe dans les .tf & le detail de l'install local)

A quoi servent pr le moment le bout de code dans provider.go ( les methode de la lib, provider 'vide')

on init le projet go ( bien préciser qu'il faut être dans le dossier terraform-provider-courses)
go mod init terraform-provider-courses
go mod tidy 

-- STEP 1 --
Nous allons commencer par créer une source de données qui va correspondre à un getAll de notre api, pour se mettre en bouche !

on crée le fichier data_source_courses.go avec tout le contenue nécessaire, ou alors pas à pas ?
On commence pas créer le Schéma de notre objet, qui correspond à celui de l'api ( on revient pas sur le détail du go)
Maintenant qu'on a défini le schéma de l'objet on peut implémenter la méthode pour requeter notre API

modifier le provider pour y ajouter la datasource, montrer que c'est bien le nom de la ressource terraform qu'on est en train de créer
/!\ Attention peut être indiquer que si le makefile plante ca peut etre lié soit au arch soit au path de terraform.d/plugin qui existe pas

Montrer le res du plan avec la liste de cours

-- STEP 2 -- Create

On vien créer le nouveau fichier data_source_cour
On crée le nouveau schéma ( id en optional pour éviter de devoir le set dans le TF mais qu'on puisse le récuperer depuis un get si besoin)

On init quand même le CRUB mais si on renvoit tout a nil, histoire de pas revenir sur des ligne de code plusieurs fois
expliquer que les Method Readcontext ect sont des fonction attendue de l objet schema fourni par terraform, ce sont les fonctons qui s execute pour chaque type d'action sur un objet TF
Expliquer  le "d." qui est la ref de la ressource actuel modifié 
peut être passer rapidement sur me détail du code go ( les call http le handle d erreur ect)
Parler du d.SetId, la on met l'id de la resources créé mais on peut imaginer n'importe quoi, TANT QUE CA RESTE UNIQUE ( par exemple pour le provider terraform, pour ton obj)
Modifier le main.tf pour y ajouter une ressoures, lancer le tf apply, aller sur le front pour le resultat

-- STEP 3 -- Get / Update

On va venir populate les deux méthodes Read & Update,
BIEN METTRE EN EVIDENCE que la last update est obligatoire côté TF, pas forcément côté API

On implémente les méthode read & update
insister encore sur le bout de code qui set le last update

Modifier le main.tf modifier la resource, lancer le tf apply, aller sur le front pour le resultat

-- STEP 4 -- DELETE

On implemente notre derniere méthode, le DeleteContext
Seul point important dans le Delete & le setId a null 

Modifier le main.tf, supprimer le code de la ressource, lancer le tf apply, et go le front

Done !

### Conseils & conclusion

Nous vennons, en quelques minutes, de donner la possibilité de consommer un service, et le tout as code.
Permettant ainsi a nos utilisateurs de pouvoir consommer notre service de manière très simple, via le code déclarative de terraform, pouvant automatiser tout son processus de validation terraform via des CI les plus utilisés. Ceci pourrait vous débloquer dans la mise en place de vos processus devops ;)

--> développer l'impatc sur le devops §( cloud privée & côté consommation)

Les possibilités que nous avont vu ici ne sont bien sur pas exaustive du tout ! On peut imaginer ajouter des paramêtre a notre provider, comme un projet en particulier, une organisation, une zone géographique que sais-je, mais aussi bien sur une couche de securité, avec des token d'api a passer en entrée du provider. ( ** peut être une suite ? mais c'est pas trop long de ce que j'ai vu)
