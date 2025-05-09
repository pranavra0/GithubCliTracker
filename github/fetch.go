package github

func FetchRepoData(repos []string, token string) ([]*RepoData, error) {
	var repoData []*RepoData
	for _, repo := range repos {
		data, err := FetchRepo(repo, token)
		if err != nil {
			return nil, err
		}
		repoData = append(repoData, data)
	}
	return repoData, nil
}
