package v1

import (
	"net/http"

	"go-postgres-microblog-openshift/internal/data"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/posts", pr.Routes())

	return r
}
