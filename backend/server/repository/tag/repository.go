package tag

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

// TODO: 絶対に S3 か GCS に移行する
// 現状はローカルに置いているので運用があまりにも大変

type Repository interface {
	GetRandomTags(contestSlug string, num int) ([]string, error)
	GetTags(contestSlug string, count int) ([]string, error)
	CreateRandomTag(contestSlug string, tags []string) error
	CreateTags(contestSlug string, tagsList [][]string) error
}

type repository struct {
	storagePath string
}

func NewRespository(storagePath string) Repository {
	return &repository{storagePath}
}

func (r *repository) GetRandomTags(contestSlug string, num int) ([]string, error) {
	b, err := os.ReadFile(filepath.Join(r.storagePath, fmt.Sprintf("storage/tags/%s/random.txt", contestSlug)))
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

func (r *repository) GetTags(contestSlug string, count int) ([]string, error) {
	b, err := os.ReadFile(filepath.Join(r.storagePath, fmt.Sprintf("storage/tags/%s/%d.txt", contestSlug, count)))
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(b)), nil
}

// random.txt will be created
func (r *repository) CreateRandomTag(contestSlug string, tags []string) error {
	contestDir := filepath.Join(r.storagePath, fmt.Sprintf("storage/tags/%s", contestSlug))
	_ = os.MkdirAll(contestDir, fs.ModeDir)
	f, err := os.Create(filepath.Join(contestDir, "random.txt"))
	if err != nil {
		return err
	}
	defer f.Close()
	return writeTags(f, tags)
}

// 1.txt, 2.txt will be created
func (r *repository) CreateTags(contestSlug string, tagsList [][]string) error {
	contestDir := filepath.Join(r.storagePath, fmt.Sprintf("storage/tags/%s", contestSlug))
	_ = os.MkdirAll(contestDir, fs.ModeDir)
	for i, tags := range tagsList {
		f, err := os.Create(filepath.Join(contestDir, fmt.Sprintf("%d.txt", i+1)))
		if err != nil {
			return err
		}
		if err := writeTags(f, tags); err != nil {
			return err
		}
		f.Close()
	}
	return nil
}

func writeTags(w io.Writer, tags []string) error {
	bufw := bufio.NewWriter(w)
	defer bufw.Flush()
	for _, tag := range tags {
		if _, err := bufw.WriteString(tag); err != nil {
			return err
		}
		if err := bufw.WriteByte('\n'); err != nil {
			return err
		}
	}
	return nil
}
