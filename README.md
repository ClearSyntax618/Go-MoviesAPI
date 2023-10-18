# Movies API with Go & Gorilla Mux
Este es mi primer proyecto utilizando Go y Gorilla Mux. Es una API simple que me sirvió para entender como funciona Go en el Backend.
Ésta API consta de cinco endpoints. Dos son `GET`, uno es `POST`, uno es `PUT` y el último es `DELETE`. 

| Métodos | URL | Funcion |
| -------- | ------- | --------------- |
| `GET` | /movies | getMovies |
| `GET` | /movies/`{id}` | getMovie |
| `POST` | /movies | createMovie |
| `PUT` | /movies/`{id}` | updateMovie |
| `DELETE` | /movies/`{id}` | deleteMovie |

Cada endpoint retorna un objeto `JSON`.

### Estructura de cada película
**Ejemplo**

```go
// No utilizaremos una base de datos. Si no que utilizaremos structs y slices
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
```
