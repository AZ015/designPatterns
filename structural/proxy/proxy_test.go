package proxy

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestUserListProxy(t *testing.T) {
	someDB := UserList{}

	rand.Seed(time.Now().Unix())

	for i := 0; i < 10000; i++ {
		n := rand.Int31()
		someDB = append(someDB, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDB,
		StackCapacity: 2,
		StackCache:    UserList{},
	}

	knownIDs := [3]int32{someDB[3].ID, someDB[4].ID, someDB[5].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])

		require.NoError(t, err)
		require.Equal(t, user.ID, knownIDs[0])
		require.Len(t, proxy.StackCache, 1)
		require.False(t, proxy.DidLastSearchUsedCache)
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])

		require.NoError(t, err)
		require.Equal(t, user.ID, knownIDs[0])
		require.Len(t, proxy.StackCache, 1)
		require.False(t, proxy.DidLastSearchUsedCache)

		user1, err := proxy.FindUser(knownIDs[0])
		require.NoError(t, err)
		proxy.FindUser(knownIDs[1])
		require.False(t, proxy.DidLastSearchUsedCache)
		proxy.FindUser(knownIDs[2])
		require.False(t, proxy.DidLastSearchUsedCache)

		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}

		require.Len(t, proxy.StackCache, 2)
	})
}
