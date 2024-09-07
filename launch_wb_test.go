//nolint:testpackage // Test it's easier with white-box testing
package neolauncher

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLaunchFailure(t *testing.T) {
	quitCh := make(chan os.Signal, 1)
	err := launch(context.Background(), quitCh, nil)
	require.ErrorIs(t, err, ErrLauncherIsNil)
}

func TestLaunch(t *testing.T) {
	t.Parallel()

	l := &MockLaunch{}
	l.On("Launch", mock.Anything).Return(nil)

	quitCh := make(chan os.Signal, 1)

	go func() {
		time.Sleep(500 * time.Millisecond)
		quitCh <- syscall.SIGTERM
	}()

	err := launch(context.Background(), quitCh, l)
	require.NoError(t, err)

	l.AssertExpectations(t)
}

type MockLaunch struct {
	mock.Mock
}

func (m *MockLaunch) Launch(ctx context.Context) error {
	args := m.Called(ctx)
	err := args.Error(0)
	if err == nil {
		<-ctx.Done()
	}
	return err
}
