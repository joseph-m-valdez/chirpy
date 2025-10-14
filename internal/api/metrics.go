package api

import (
	"fmt"
	"net/http"
)

func (a *API) HandlerMetrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	serverHits := fmt.Sprintf(`
<html>

<body>
	<h1>Welcome, Chirpy Admin</h1>
	<p>Chirpy has been visited %d times!</p>
</body>

</html>
	`, a.FileServerHits.Load())

	w.Write([]byte(serverHits))
}
