package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	client "gogoapps-go/pkg/client"
	config2 "gogoapps-go/pkg/config"
	"net/http"
)

type api struct {
	config config2.Config
	router *mux.Router
	client client.PictureClient
}

func New(cfg config2.Config, client client.PictureClient) (*api, error) {
	return &api{
		config: cfg,
		router: mux.NewRouter(),
		client: client,
	}, nil
}

func (api *api) SetServer() error {
	if err := api.SetRoutes(); err != nil {
		return err
	}

	if err := http.ListenAndServe(":"+api.config.Server.Port, api.router); err != nil {
		return fmt.Errorf("something went wrong: %w", err)
	}
	return nil
}

func (api *api) SetRoutes() error {
	pictureStartDate := api.config.Picture.StartDate
	pictureEndDate := api.config.Picture.EndDate

	picturesEndpoint := api.router.HandleFunc("/pictures", api.PicturesHandler()).Methods("GET").Queries(pictureStartDate,
		"{"+pictureStartDate+"}").Queries(pictureEndDate,
		"{"+pictureEndDate+"}")

	if picturesEndpoint.GetError() != nil {
		return picturesEndpoint.GetError()
	}
	return nil
}

func (api *api) PicturesHandler() http.HandlerFunc {
	type response struct {
		Result   []string `json:"urls,omitempty"`
		ErrorMsg string   `json:"error,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var response response
		vars := mux.Vars(r)

		result, err := api.client.GetPicture(vars)
		if err != nil {
			response.ErrorMsg = err.Error()
			_ = json.NewEncoder(w).Encode(response)
			return
		}

		response.Result = result
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			response.ErrorMsg = err.Error()
			return
		}
	}
}
