package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Client struct {
	BaseUrl string
}

type MachineData struct {
	State int `json:"state"`
}

type Balancer struct {
	Id 			  	   int   `json:"id"`
	UsedMachines  	   []int `json:"usedmachines"`
	TotalMachinesCount int   `json:"totalmachinescount"`
}

func (c *Client) ListBalancers(ctx context.Context) ([]Balancer, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseUrl+"/balancers", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
	defer res.Body.Close()

	var result []Balancer
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) ChangeMachineStatus(ctx context.Context, id int, state bool) error {
	var data MachineData
	if state {
		data.State = 1
	} else {
		data.State = 0
	}
	bb := new(bytes.Buffer)
	err := json.NewEncoder(bb).Encode(&data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, c.BaseUrl+"/machines/"+strconv.Itoa(id), bb)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
	defer res.Body.Close()

	return nil
}