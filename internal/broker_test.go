package internal_test

import (
	"message_broker/internal/channels"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannels(t *testing.T) {
	ch := channels.NewMessageChannel(5)

	testValues := []string{"Hello", "World", "Apple", "Banana"}
	result := make([]string, 0, 4)
	
	for _, el := range testValues{
		ch.Send([]byte(el))
	}

	for i := 0; i < 4; i++ {
		value, _ := ch.Receive()
		result = append(result, string(value))
	}

	assert.Equal(t, len(testValues), len(result))
	for i := 0; i < 4; i++ {
		assert.Equal(t, testValues[i], result[i])
	}
}