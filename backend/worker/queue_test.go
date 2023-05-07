package worker_test

import (
	"testing"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/worker"
	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	var alice, bob, charlie int
	q := &worker.Queue[int]{}
	q.Push(&alice)
	q.Push(&bob)
	q.Push(&charlie)
	assert.Equal(t, &alice, q.Pop())
	assert.Equal(t, &bob, q.Pop())
	assert.Equal(t, &charlie, q.Pop())
	assert.Nil(t, q.Pop())
}
