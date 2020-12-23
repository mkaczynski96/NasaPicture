package client

import (
	"fmt"
	"gogoapps-go/pkg/client/nasa"
	"gogoapps-go/pkg/config"
	"runtime"
	"sync"
)

type PictureClient interface {
	GetPicture(vars map[string]string) ([]string, error)
}

type Client struct {
	cfg config.Config
}

func NewClient(cfg config.Config) (*Client, error) {
	return &Client{cfg: cfg}, nil
}

func (client *Client) GetPicture(vars map[string]string) ([]string, error) {
	var result []string

	// Get date parameters from path variables
	startDate, endDate, err := StartEndDate(vars)
	if err != nil {
		return nil, err
	}

	// Validation
	if err := ValidParameters(startDate, endDate); err != nil {
		return nil, err
	}

	// Get date range
	dates, err := DateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Request for picture
	baseUrl := fmt.Sprintf(client.cfg.External.Url + client.cfg.External.ApiKey + client.cfg.External.DateParameter)
	pictures, err := client.RequestPictures(baseUrl, dates)
	if err != nil {
		return nil, err
	}

	// Append picture url to result
	for _, picture := range pictures {
		nasaPicture, err := nasa.UnmarshallPicture(picture)
		if err != nil {
			return nil, err
		}
		if nasaPicture != nil {
			if len(nasaPicture.Url) != 0 {
				result = append(result, nasaPicture.Url)
			}
		}
	}

	return result, nil
}

func (client *Client) RequestPictures(baseUrl string, dates []string) (map[int][]byte, error) {
	pictureMap := make(map[int][]byte, len(dates))

	maxRequests := client.cfg.Server.MaxRequests
	maxConcurrent := client.cfg.Server.ConcurrentRequests
	runtime.GOMAXPROCS(runtime.NumCPU())
	rc := make(chan Resp)
	sem := make(chan bool, maxConcurrent)

	wg := sync.WaitGroup{}
	for i, date := range dates {
		wg.Add(1)
		go func(i int, date string) {
			urlWithDate := fmt.Sprintf(baseUrl + date)
			go MakeGetRequest(urlWithDate, maxRequests, rc, sem)
			response := CheckResponses(rc)
			pictureMap[i] = response
			wg.Done()
		}(i, date)
	}
	wg.Wait()
	return pictureMap, nil
}
