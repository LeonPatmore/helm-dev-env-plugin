// func getCiYaml(repo string, client github.Client) (*CIConfig, error) {
// 	ctx := context.Background()
// 	fileContent, _, _, err := client.Repositories.GetContents(ctx, "nexmoinc", repo, "ci.yaml", &github.RepositoryContentGetOptions{Ref: "master"})
// 	if err != nil {
// 		return nil, err
// 	}
// 	yamlStr, err := fileContent.GetContent()
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(yamlStr)
// 	config := &CIConfig{}
// 	err = yaml.Unmarshal([]byte(yamlStr), config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return config, nil
// }

package main
