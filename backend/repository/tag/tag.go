package tag

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

// TODO: 絶対に S3 か GCS に移行する
// 現状はローカルに置いているので運用があまりにも大変

func GetRandomTags(contestID, num int) ([]string, error) {
	b, err := os.ReadFile(fmt.Sprintf("./storage/tags/%d/random.txt", contestID))
	if err != nil {
		return nil, err
	}
	tags := strings.Fields(string(b))
	shuffledTags := lo.Shuffle(tags)
	if len(shuffledTags) < num {
		return shuffledTags, nil
	} else {
		return shuffledTags[:num], nil
	}
}

func GetTags(contestID, count int) ([]string, error) {
	b, err := os.ReadFile(fmt.Sprintf("./storage/tags/%d/%d.txt", contestID, count))
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(b)), nil
}
