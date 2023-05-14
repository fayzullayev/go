package main

import (
	"github.com/fayzullayev/snippetbox/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		app.notFound(writer)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignUp))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignUpPost))
	router.Handler(http.MethodGet, "/user/signIn", dynamic.ThenFunc(app.userSignIn))
	router.Handler(http.MethodPost, "/user/signIn", dynamic.ThenFunc(app.userSignInPost))

	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogOut))

	//router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	//router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	//router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	//router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))
	//
	//router.Handler(http.MethodGet, "/user/signUp", dynamic.ThenFunc(app.userSignUp))
	//router.Handler(http.MethodGet, "/user/signIn", dynamic.ThenFunc(app.userSignIn))
	//
	//router.Handler(http.MethodPost, "/user/signUp", dynamic.ThenFunc(app.userSignUpPost))
	//router.Handler(http.MethodPost, "/user/signIn", dynamic.ThenFunc(app.userSignInPost))
	//
	//router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.userLogOut))

	//router.HandlerFunc(http.MethodGet, "/", app.home)
	//router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	//router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	//router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	//router.HandlerFunc(http.MethodGet, "/", app.home)
	//mux := http.NewServeMux()
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	//
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//
	//mux.HandleFunc("/", app.home)
	//mux.HandleFunc("/snippet/view", app.snippetView)
	//mux.HandleFunc("/snippet/create", app.snippetCreate)
	//

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
