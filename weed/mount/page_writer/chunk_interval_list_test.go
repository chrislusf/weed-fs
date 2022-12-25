package page_writer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PageChunkWrittenIntervalList(t *testing.T) {
	list := newChunkWrittenIntervalList()

	assert.Equal(t, 0, list.size(), "empty list")

	list.MarkWritten(0, 5, 1)
	assert.Equal(t, 1, list.size(), "one interval")

	list.MarkWritten(0, 5, 2)
	assert.Equal(t, 1, list.size(), "duplicated interval2")

	list.MarkWritten(95, 100, 3)
	assert.Equal(t, 2, list.size(), "two intervals")

	list.MarkWritten(50, 60, 4)
	assert.Equal(t, 3, list.size(), "three intervals")

	list.MarkWritten(50, 55, 5)
	assert.Equal(t, 4, list.size(), "three intervals merge")

	list.MarkWritten(40, 50, 6)
	assert.Equal(t, 5, list.size(), "three intervals grow forward")

	list.MarkWritten(50, 65, 7)
	assert.Equal(t, 4, list.size(), "three intervals grow backward")

	list.MarkWritten(70, 80, 8)
	assert.Equal(t, 5, list.size(), "four intervals")

	list.MarkWritten(60, 70, 9)
	assert.Equal(t, 6, list.size(), "three intervals merged")

	list.MarkWritten(59, 71, 10)
	assert.Equal(t, 6, list.size(), "covered three intervals")

	list.MarkWritten(5, 59, 11)
	assert.Equal(t, 5, list.size(), "covered two intervals")

	list.MarkWritten(70, 99, 12)
	assert.Equal(t, 5, list.size(), "covered one intervals")

}
