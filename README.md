# Movies API with Go & Gorilla Mux
Este es mi primer proyecto utilizando Go y Gorilla Mux. Es una API simple que me sirvió para entender como funciona Go en el Backend.
Ésta API consta de cinco endpoints. Dos son `GET`, uno es `POST`, uno es `PUT` y el último es `DELETE`. 

| Métodos | URL |
| -------- | ------- |
| `GET` | /movies |
| `GET` | /movies/`{id}` |
| `POST` | /movies |
| `PUT` | /movies/`{id}` |
| `DELETE` | /movies/`{id}` |

Cada endpoint retorna un objeto `JSON`.

**Ejemplo no real del retorno de la API**

```json
{
  "name": "Jurassic Park",
  "year": 1993,
  "runtime": 127,
  "categories": [
    "adventure",
    "thriller",
    "sci-fi"
  ],
  "release-date": "1993-06-11",
  "director": "Steven Spielberg",
  "writer": [
    "Michael Crichton",
    "David Koepp"
  ],
  "actors": [
    "Sam Neill",
    "Laura Dern",
    "Jeff Goldblum"
  ],
  "storyline": "Huge advancements in scientific technology have enabled a mogul ... critical security systems are shut down and it now becomes a race for survival with dinosaurs roaming freely over the island."
}
```
