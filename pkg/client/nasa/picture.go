package nasa

import (
	"encoding/json"
	"fmt"
)

type PictureNasa struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

func UnmarshallPicture(unmarshalledPicture []byte) (nasaPicture *PictureNasa, err error) {
	err = json.Unmarshal(unmarshalledPicture, &nasaPicture)
	if err != nil {
		return nil, fmt.Errorf("smth went wrong")
	}
	return nasaPicture, nil
}
