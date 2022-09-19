package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kmtym1998/hasuraenv/internal/httpc"
	"github.com/kmtym1998/hasuraenv/internal/model"
)

func ListHasuraReleases() ([]model.GitHubRelease, error) {
	httpResp, err := httpc.SendRequest(
		http.MethodGet,
		"https://api.github.com/repos/hasura/graphql-engine/releases",
		nil,
		map[string]string{
			"Content-Type": "application/json",
		},
	)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	b, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var releases []model.GitHubRelease
	if err := json.Unmarshal(b, &releases); err != nil {
		return nil, err
	}

	return releases, nil
}
