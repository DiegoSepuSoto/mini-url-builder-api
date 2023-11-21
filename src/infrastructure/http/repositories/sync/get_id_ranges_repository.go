package sync

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/DiegoSepuSoto/mini-url-builder-api/src/shared"
)

type IDRangesResponse struct {
	RangeIDs string `json:"rangeIDs"`
}

func (r *syncRepository) GetIDRanges() (string, error) {
	url := fmt.Sprintf("%s/%s", os.Getenv("SYNC_SERVICE_HOST"), "range")

	resp, err := r.HTTPClient.Get(url)
	if err != nil {
		return "", shared.BuildError(http.StatusInternalServerError,
			shared.SyncCommunicationError,
			err.Error(),
			"syncRepository")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", shared.BuildError(http.StatusInternalServerError,
			"SYNC_RESP_ERROR",
			err.Error(),
			"SyncRepository")
	}

	data := IDRangesResponse{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return "", shared.BuildError(http.StatusInternalServerError,
			"SYNC_RESP_ERROR",
			err.Error(),
			"SyncRepository")
	}

	return data.RangeIDs, nil
}
