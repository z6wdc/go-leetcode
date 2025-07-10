package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/z6wdc/go-leetcode/internal/entity"
	"io"
	"net/http"
	"strconv"
)

const graphqlURL = "https://leetcode.com/graphql"

type questionGraphQLResponse struct {
	Data struct {
		Question struct {
			QuestionID string `json:"questionId"`
			Title      string `json:"title"`
			TitleSlug  string `json:"titleSlug"`
			Content    string `json:"content"`
			Difficulty string `json:"difficulty"`
			TopicTags  []struct {
				Name string `json:"name"`
			} `json:"topicTags"`
		} `json:"question"`
	} `json:"data"`
}

// FetchQuestionBySlug fetches question data from leetcode by titleSlug
func FetchQuestionBySlug(slug string) (*entity.Question, error) {
	payload := map[string]interface{}{
		"operationName": "questionData",
		"variables": map[string]string{
			"titleSlug": slug,
		},
		"query": `
			query questionData($titleSlug: String!) {
				question(titleSlug: $titleSlug) {
					questionId
					title
					titleSlug
					content
					difficulty
					topicTags {
						name
					}
				}
			}
		`,
	}
	bodyBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", graphqlURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("LeetCode API returned status %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	var gqlResp questionGraphQLResponse
	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return nil, err
	}

	q := gqlResp.Data.Question
	id, _ := strconv.Atoi(q.QuestionID)

	// Convert HTML to Markdown
	mdContent, err := htmltomarkdown.ConvertString(q.Content)
	if err != nil {
		return nil, err
	}

	var tags []string
	for _, tag := range q.TopicTags {
		tags = append(tags, tag.Name)
	}

	return &entity.Question{
		ID:         id,
		Title:      q.Title,
		Slug:       q.TitleSlug,
		Content:    mdContent,
		Difficulty: q.Difficulty,
		Tags:       tags,
	}, nil
}
