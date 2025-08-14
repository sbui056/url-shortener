package shortener

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const UserID = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestGenerateShortLink_UniqueGeneration(t *testing.T) {
	link1 := "https://www.example.com/page1"
	link2 := "https://www.example.com/page2"

	short1 := GenerateShortLink(link1, UserID)
	short2 := GenerateShortLink(link2, UserID)

	assert.NotEmpty(t, short1, "Short link for link1 should not be empty")
	assert.NotEmpty(t, short2, "Short link for link2 should not be empty")
	assert.NotEqual(t, short1, short2, "Different URLs should generate different short links")
}

func TestGenerateShortLink_Consistency(t *testing.T) {
	link := "https://www.example.com/page"
	short1 := GenerateShortLink(link, UserID)
	short2 := GenerateShortLink(link, UserID)

	assert.Equal(t, short1, short2, "Same URL and user ID should generate the same short link")
}

func TestGenerateShortLink_DifferentUsers(t *testing.T) {
	link := "https://www.example.com/page"
	user1 := "user-123"
	user2 := "user-456"

	short1 := GenerateShortLink(link, user1)
	short2 := GenerateShortLink(link, user2)

	assert.NotEqual(t, short1, short2, "Same URL with different user IDs should generate different short links")
}
