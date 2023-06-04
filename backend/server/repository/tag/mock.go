package tag

// TODO: 絶対に S3 か GCS に移行する
// 現状はローカルに置いているので運用があまりにも大変

type mockRepository struct {
	getRandomTags func(contestID, num int) ([]string, error)
	getTags       func(contestID, count int) ([]string, error)
}

func MockRepository(
	getRandomTags func(contestID, num int) ([]string, error),
	getTags func(contestID, count int) ([]string, error),
) Repository {
	return &mockRepository{getRandomTags, getTags}
}

func (r *mockRepository) GetRandomTags(contestID, num int) ([]string, error) {
	if r.getRandomTags != nil {
		return r.getRandomTags(contestID, num)
	}
	return []string{}, nil
}

func (r *mockRepository) GetTags(contestID, count int) ([]string, error) {
	if r.getTags != nil {
		return r.getTags(contestID, count)
	}
	return []string{}, nil
}
