package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

type LocationExtractorFromRequest struct {
	Request *http.Request
}

type loopkupResponse struct {
	Success string  `json:"success"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

func (re *LocationExtractorFromRequest) GetLocation() (Location, Location, error) {
	emptyLocation := Location{}
	ip, _, err := net.SplitHostPort(re.Request.RemoteAddr)
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to get host:%w", err)
	}

	url := fmt.Sprintf("http://ip-api.com/json/%v?fields=status,lat,lon", ip)
	fmt.Println(url, ip)
	res, err := http.Get(url)
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("GeoLookup failed:%w", err)
	}
	if res.StatusCode != http.StatusOK {
		return emptyLocation, emptyLocation, fmt.Errorf("GeoLookup failed with status code:%v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to read geoLookup response:%w", err)
	}

	var resp loopkupResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return emptyLocation, emptyLocation, fmt.Errorf("Failed to parse geolookup response:%w", err)
	}
	if resp.Success != "success" {
		return emptyLocation, emptyLocation, fmt.Errorf("GeoLookup failed")
	}

	lat := Location{}
	lat.FromFloat(resp.Lat, true)

	lon := Location{}
	lon.FromFloat(resp.Lat, false)

	return lat, lon, nil
}
