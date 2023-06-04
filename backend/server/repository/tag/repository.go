package tag

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

// TODO: 絶対に S3 か GCS に移行する
// 現状はローカルに置いているので運用があまりにも大変

type Repository interface {
	GetRandomTags(contestID, num int) ([]string, error)
	GetTags(contestID, count int) ([]string, error)
}

type repository struct {
	storagePath string
}

func NewRespository(storagePath string) Repository {
	return &repository{storagePath}
}

func (r *repository) GetRandomTags(contestID, num int) ([]string, error) {
	b, err := os.ReadFile(filepath.Join(r.storagePath, fmt.Sprintf("storage/tags/%d/random.txt", contestID)))
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

func (r *repository) GetTags(contestID, count int) ([]string, error) {
	b, err := os.ReadFile(filepath.Join(r.storagePath, fmt.Sprintf("storage/tags/%d/%d.txt", contestID, count)))
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(b)), nil
}
