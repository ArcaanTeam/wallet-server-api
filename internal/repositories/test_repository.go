package repositories

type testRepo struct{}

func NewTestRepo() *testRepo {
	return &testRepo{}
}

func (repo *testRepo) Ping() string {
	return "pong"
}
