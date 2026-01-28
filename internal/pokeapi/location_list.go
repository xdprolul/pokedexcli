package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

func(c *Client) ListLocations(pageURL *string) (ResShallowLocations, error) {
	url:=baseURL+"/location-area"
	if pageURL!=nil {
		url=*pageURL
	}

	req,err:=http.NewRequest("GET",url,nil)
	if err!=nil {
		return ResShallowLocations{},err
	}

	res,err:=c.httpClient.Do(req)
	if err!=nil {
		return ResShallowLocations{},err
	}
	defer res.Body.Close()

	data,err:=io.ReadAll(res.Body)
	if err!=nil {
		return ResShallowLocations{},err
	}

	locationRes:=ResShallowLocations{}
	err = json.Unmarshal(data,&locationRes)
	if err!=nil {
		return ResShallowLocations{},err
	}

	return locationRes,nil
}
