package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kmtym1998/hasuraenv/internal/httpc"
	"github.com/kmtym1998/hasuraenv/internal/model"
)

// https://docs.github.com/ja/rest/releases/releases#list-releases

const GITHUB_ENDPOINT_GET_RELEASE_BY_TAG_NAME string = "https://api.github.com/repos/hasura/graphql-engine/releases/tags/%s"

func GetReleaseByTagName(tag string) (*model.GitHubRelease, error) {
	u, _ := url.Parse(fmt.Sprintf(GITHUB_ENDPOINT_GET_RELEASE_BY_TAG_NAME, tag))

	httpResp, err := httpc.SendRequest(
		http.MethodGet,
		u.String(),
		nil,
		map[string]string{
			"Content-Type": "application/json",
		},
	)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	b, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var result *model.GitHubRelease
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return result, nil
}
