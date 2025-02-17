package prover_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-service/eigenda/encoding"
	"github.com/ethereum-optimism/optimism/op-service/eigenda/encoding/kzg"
	"github.com/ethereum-optimism/optimism/op-service/eigenda/encoding/kzg/prover"
)

func TestNewSRSTable_PreComputeWorks(t *testing.T) {

	kzgConfig.CacheDir = "./data/SRSTable"
	params := encoding.ParamsFromSysPar(numSys, numPar, uint64(len(gettysburgAddressBytes)))
	require.NotNil(t, params)

	s1, err := kzg.ReadG1Points(kzgConfig.G1Path, kzgConfig.SRSOrder, kzgConfig.NumWorker)
	require.Nil(t, err)
	require.NotNil(t, s1)

	_, err = kzg.ReadG2Points(kzgConfig.G2Path, kzgConfig.SRSOrder, kzgConfig.NumWorker)
	require.Nil(t, err)

	subTable1, err := prover.NewSRSTable(kzgConfig.CacheDir, s1, kzgConfig.NumWorker)
	require.Nil(t, err)
	require.NotNil(t, subTable1)

	fftPoints1, err := subTable1.GetSubTables(params.NumChunks, params.ChunkLength)
	require.Nil(t, err)
	require.NotNil(t, fftPoints1)

	subTable2, err := prover.NewSRSTable(kzgConfig.CacheDir, s1, kzgConfig.NumWorker)
	require.Nil(t, err)
	require.NotNil(t, subTable2)

	fftPoints2, err := subTable2.GetSubTables(params.NumChunks, params.ChunkLength)
	require.Nil(t, err)
	require.NotNil(t, fftPoints2)

	// Result of non precomputed GetSubTables should equal precomputed GetSubTables
	assert.Equal(t, fftPoints1, fftPoints2)
}
