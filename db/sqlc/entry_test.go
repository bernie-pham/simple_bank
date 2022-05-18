package db

import (
	"context"
	"testing"
	"time"

	"github.com/bernie-pham/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestReadEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	arg := ListEntryParams{
		Limit:  5,
		Offset: 0,
	}
	entries, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, int(arg.Limit))

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		// require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
