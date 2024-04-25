package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.session.Exists(r, "authenticatedUserID")
}

func (app *application) whoIsThis(r *http.Request) string {
	if app.session.Exists(r, "authenticatedUserID") {
		currentId := app.session.Get(r, "authenticatedUserID").(int)
		return app.client.GetUserRoleById(currentId) // todo create users model
		// todo
	}
	return "guest"
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	//td.Flash = app.session.PopString(r, "flash")
	//td.CommentFlash = app.session.PopString(r, "commentFlash")

	td.IsAuthenticated = app.isAuthenticated(r)
	td.Role = app.whoIsThis(r)
	return td
}

//func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
//	ts, ok := app.templateCache[name]
//	if !ok {
//		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
//		return
//	}
//
//	buf := new(bytes.Buffer)
//
//	// Execute the template set, passing the dynamic data with the current
//	// year injected.
//	err := ts.Execute(buf, app.addDefaultData(td, r))
//	if err != nil {
//		app.serverError(w, err)
//		return
//	}
//
//	buf.WriteTo(w)
//}

type envelope map[string]interface{}

// Retrieve the "id" URL parameter from the current request context, then convert it to
// an integer and return it. If the operation isn't successful, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// todo: go get github.com/julienschmidt/httprouter
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

// in my version of go there is no type as 'any', and instead of it I used interface{},
// cuz Marshal actually accepts it as a parameter and map is implementing interface.
// on your side data interface{} must be data any if you are using go version 1.18 or newer
// any is a type alias of interface
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	//adding additional headers if there are any to be added
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Adding Content-Type and status code to header and response as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		if errors.As(err, &syntaxError) {
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		} else if errors.As(err, &unmarshalTypeError) {
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", unmarshalTypeError.Offset)
		} else if errors.As(err, &invalidUnmarshalError) {
			panic(err) //If our program reaches a point where it cannot be recovered due to some major errors

		} else if errors.Is(err, io.ErrUnexpectedEOF) {
			return errors.New("body contains badly-formed JSON")
		} else if errors.Is(err, io.EOF) {
			return errors.New("body must not be empty")
		} else {
			return err
		}
	}

	return nil
}
