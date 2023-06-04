package tag

// TODO: 絶対に S3 か GCS に移行する
// 現状はローカルに置いているので運用があまりにも大変

type mockRepository struct {
	getRandomTags func(contestSlug string, num int) ([]string, error)
	getTags       func(contestSlug string, count int) ([]string, error)
}

func MockRepository(
	getRandomTags func(contestSlug string, num int) ([]string, error),
	getTags func(contestSlug string, count int) ([]string, error),
) Repository {
	return &mockRepository{getRandomTags, getTags}
}

func (r *mockRepository) GetRandomTags(contestSlug string, num int) ([]string, error) {
	if r.getRandomTags != nil {
		return r.getRandomTags(contestSlug, num)
	}
	return []string{}, nil
}

func (r *mockRepository) GetTags(contestSlug string, count int) ([]string, error) {
	if r.getTags != nil {
		return r.getTags(contestSlug, count)
	}
	return []string{}, nil
}
