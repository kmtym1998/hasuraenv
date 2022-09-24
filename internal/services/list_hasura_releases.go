package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/kmtym1998/hasuraenv/internal/httpc"
	"github.com/kmtym1998/hasuraenv/internal/model"
)

// https://docs.github.com/ja/rest/releases/releases#list-releases

const GITHUB_ENDPOINT_LIST_RELEASES string = "https://api.github.com/repos/hasura/graphql-engine/releases"

func ListHasuraReleases(limit int) ([]model.GitHubRelease, error) {
	over100 := limit > 100

	u, _ := url.Parse(GITHUB_ENDPOINT_LIST_RELEASES)
	queryParams := u.Query()

	var releases []model.GitHubRelease
	for i := 1; len(releases) < limit; i++ {
		queryParams.Set("page", strconv.Itoa(i))
		if over100 {
			queryParams.Set("per_page", "100")
		} else {
			queryParams.Set("per_page", strconv.Itoa(limit))
		}

		u.RawQuery = queryParams.Encode()

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

		b, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}

		var resp []model.GitHubRelease
		if err := json.Unmarshal(b, &resp); err != nil {
			return nil, err
		}

		if len(resp) == 0 {
			break
		}

		releases = append(releases, resp...)
	}

	var results []model.GitHubRelease
	for i := 0; i < limit; i++ {
		if i >= len(releases) {
			break
		}

		results = append(results, releases[i])
	}

	return results, nil
}
