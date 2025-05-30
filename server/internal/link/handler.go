package link

import (
	"fmt"
	"net/http"
	"strconv"

	"go/adv-demo/pkg/middleware"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/response"

	"gorm.io/gorm"
)

type (
	LinkHandlerDeps struct {
		LinkRepository *LinkRepository
	}
	LinkHandler struct {
		LinkRepository *LinkRepository
	}
)

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	// fmt.Println("NewLinkHandler")
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update()))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.Goto())
	// router.HandleFunc("GET /link/{url}", handler.GetByUrl)
	// router.HandleFunc("POST /auth/register", handler.Register)
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("CreateHandler")
		// handler.LinkRepository.Create(link * Link)
		body, err := request.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("UpdateHandler")
		body, err := request.HandleBody[LinkUpdateRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, *link, http.StatusCreated)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// id := req.PathValue("id")
		// fmt.Println(id)
		// response.Json(w, id, http.StatusOK)
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Вариант с поиском по id
		// _, err = handler.LinkRepository.GetByID(uint(id))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusNotFound)
		// 	return
		// }
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, nil, http.StatusOK)
	}
}

func (handler *LinkHandler) Goto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		fmt.Println("Hash = " + hash)
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Println("link.Url = " + link.Url)
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

// func (handler *LinkHandler) GetByUrl(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("GetByUrl")

// 	url := req.PathValue("url")
// 	fmt.Println("url = " + url)
// 	link, err := handler.LinkRepository.GetByUrl(url)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	fmt.Println("link.Url = " + link.Url)
// 	http.Redirect(w, req, link.Url, http.StatusTemporaryRedirect)
// }
