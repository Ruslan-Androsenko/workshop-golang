package first

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	repo := Repo{}
	repo.init()

	repo.add(1)
	repo.add(2)
	repo.add(3)

	repo.add(4)

	repo.add(1)
	repo.add(2)
	repo.add(3)

	repo.add(1)
	repo.add(2)
	repo.add(3)

	repo.add(5)

	repo.delete(3)
	repo.delete(3)
	unique := repo.get_unique()
	repo.delete(3)

	require.Greater(t, unique, 0)
	require.NoError(t, nil)
}
