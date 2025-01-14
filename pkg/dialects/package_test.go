//autogenerated:yes
//nolint:golint
package dialects

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	"github.com/aler9/gomavlib/pkg/dialects/asluav"
	"github.com/aler9/gomavlib/pkg/dialects/common"
	"github.com/aler9/gomavlib/pkg/dialects/icarous"
	"github.com/aler9/gomavlib/pkg/dialects/matrixpilot"
	"github.com/aler9/gomavlib/pkg/dialects/minimal"
	"github.com/aler9/gomavlib/pkg/dialects/paparazzi"
	"github.com/aler9/gomavlib/pkg/dialects/pythonarraytest"
	"github.com/aler9/gomavlib/pkg/dialects/standard"
	"github.com/aler9/gomavlib/pkg/dialects/test"
	"github.com/aler9/gomavlib/pkg/dialects/ualberta"
	"github.com/aler9/gomavlib/pkg/dialects/uavionix"
)

func TestDialects(t *testing.T) {
	func() {
		_, err := dialect.NewDecEncoder(asluav.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(ardupilotmega.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(common.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(icarous.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(matrixpilot.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(minimal.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(paparazzi.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(pythonarraytest.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(standard.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(test.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(uavionix.Dialect)
		require.NoError(t, err)
	}()
	func() {
		_, err := dialect.NewDecEncoder(ualberta.Dialect)
		require.NoError(t, err)
	}()
}
