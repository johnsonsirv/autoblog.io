package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/johnsonsirv/auto-blog/models"
)

type NewsController struct {
	IDPattern *regexp.Regexp
}

func (nc NewsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/news" {
		switch r.Method {
		case http.MethodGet:
			nc.getAll(w, r)
		case http.MethodPost:
			nc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := nc.IDPattern.FindStringSubmatch(r.URL.Path)
		
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id := matches[1]

		switch r.Method {
		case http.MethodGet:
			nc.get(id, w)
		case http.MethodDelete:
			nc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// Constructor Function = OOP concept
func NewNewsController() *NewsController {
	return &NewsController{
		IDPattern: regexp.MustCompile(`/news/([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})/?`),
	}
}


func (nc *NewsController) getAll(w http.ResponseWriter, r *http.Request)  {
	encodeResponseAsJSON(models.GetNews() , w)
}

func (nc *NewsController) get(id string, w http.ResponseWriter)  {
	n, err := models.GetNewsById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeResponseAsJSON(n, w)
}

func (nc *NewsController) post(w http.ResponseWriter, r *http.Request)  {
	nI, err := nc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse request object"))
		return
	}
	
	n, err := models.AddNewsItem(nI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(n, w)

}

func (nc *NewsController) delete(id string, w http.ResponseWriter)  {
	err := models.RemoveNewsById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (nc *NewsController) parseRequest(r *http.Request)(models.News, error)  {
	dec := json.NewDecoder(r.Body)

	var n models.News
	err := dec.Decode(&n)
	if err != nil {
		return models.News{}, err
	}

	return n, nil
}