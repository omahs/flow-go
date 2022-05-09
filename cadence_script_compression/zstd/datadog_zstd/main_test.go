package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
	"time"

	"github.com/DataDog/zstd"
	csc "github.com/onflow/flow-go/cadence_script_compression"
)

const (
	contracsDir = "../../contracts/mainnet"
	dictionary  = "../MAINNET_DICT"
)

func TestCompressVariableSize(t *testing.T) {
	// get a sample of contracts with variable size compared to rest of the contracts
	contractNames := []string{
		"504cbf9620ea384e.SomePlaceCounter.cdc",          // 251
		"26f07529cfd446c3.PonsCertificationContract.cdc", // 534
		"ae3baa0d314e546b.DigiyoAdminReceiver.cdc",       // 837
		"01ab36aaf654a13e.RaribleFee.cdc",                // 1090
		"ac98da57ce4dd4ef.MessageBoard.cdc",              // 1353
		"923f5c3c39a7649d.ListedTokens.cdc",              // 2872
		"097bafa4e0b48eef.CharityNFT.cdc",                // 3914
		"97165bb37258eec7.SplinterlandsItem.cdc",         // 11702
		"329feb3ab062d289.AmericanAirlines_NFT.cdc",      // 23478
		"b715b81853fef53f.AllDay.cdc",                    // 26205
		"0b2a3299cc857e29.TopShot.cdc",                   // 44575
		"2817a7313c740232.TenantService.cdc",             // 60130
		"3885d9d426d2ef5c.SpaceModel.cdc",                // 602177
		"3885d9d426d2ef5c.EndModel.cdc",                  // 405918
		"3885d9d426d2ef5c.Model.cdc",                     // 703397
	}

	contracts := csc.ReadContracts(contracsDir)

	for _, name := range contractNames {
		c := contracts[name]
		start := time.Now()

		dst := make([]byte, 0)
		dst, err := zstd.Compress(dst, c.Data)
		if err != nil {
			t.Fatal(err)
		}

		compressionSpeed := csc.CompressionSpeed(float64(len(c.Data)), start)
		ratio := csc.CompressionRatio(float64(len(c.Data)), float64(len(dst)))

		start = time.Now()
		_, err = zstd.Decompress(nil, dst)
		if err != nil {
			t.Fatal(err)
		}

		_ = time.Since(start).Microseconds()
		decompressionSpeed := csc.CompressionSpeed(float64(len(c.Data)), start)

		s := fmt.Sprintf("contract name: %s\norig size/compressed size: %d/%d\ncompression speed: %.2f MB/s\ncompression ratio: %.2f\ndecompression speed: %.2f MB/s", c.Name, len(c.Data), len(dst), compressionSpeed, ratio, decompressionSpeed)
		t.Log(s)
	}
}

func TestCompressAllContracts(t *testing.T) {
	contracts := csc.ReadContracts(contracsDir)

	sumMbPerSec := 0.00
	sumRatio := 0.00
	for _, c := range contracts {
		start := time.Now()

		dst := make([]byte, 0)
		dst, err := zstd.Compress(dst, c.Data)
		if err != nil {
			t.Fatal(err)
		}

		mbpersec := csc.CompressionSpeed(float64(len(c.Data)), start)
		sumMbPerSec = sumMbPerSec + mbpersec

		ratio := csc.CompressionRatio(float64(len(c.Data)), float64(len(dst)))
		sumRatio = sumRatio + ratio

		//t.Logf("Name: %s | UnCompressed Size: %d | Compressed Size: %d | Speed: %.2f MB/s | Ratio: %.2f", c.Name, c.Size, len(dst), mbpersec, ratio)
	}

	avgRatio := sumRatio / float64(len(contracts))
	avgSpeed := sumMbPerSec / float64(len(contracts))

	t.Logf("Average Compression Ratio: %.2f , Average Compression Speed: %.2f MB/s", avgRatio, avgSpeed)
}

func TestCompressAllWithDictBestCompression(t *testing.T) {
	contracts := csc.ReadContracts(contracsDir)

	sumMbPerSec := 0.00
	sumRatio := 0.00

	dict, err := ioutil.ReadFile(dictionary)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range contracts {
		start := time.Now()

		dst := make([]byte, 0)
		w := bytes.NewBuffer(dst)

		writer := zstd.NewWriterLevelDict(w, zstd.BestCompression, dict)
		r := bytes.NewReader(c.Data)

		_, _ = io.Copy(writer, r)
		_ = writer.Close() // Make sure the writer is closed

		dst = w.Bytes()

		mbpersec := csc.CompressionSpeed(float64(len(c.Data)), start)
		sumMbPerSec = sumMbPerSec + mbpersec

		ratio := csc.CompressionRatio(float64(len(c.Data)), float64(len(dst)))
		sumRatio = sumRatio + ratio

		//t.Logf("Name: %s | UnCompressed Size: %d | Compressed Size: %d | Speed: %.2f MB/s | Ratio: %.2f", c.Name, c.Size, len(dst), mbpersec, ratio)
	}

	avgRatio := sumRatio / float64(len(contracts))
	avgSpeed := sumMbPerSec / float64(len(contracts))

	t.Logf("Average Compression Ratio: %.2f , Average Compression Speed: %.2f MB/s", avgRatio, avgSpeed)
}

func TestCompressAllWithDictBestSpeed(t *testing.T) {
	contracts := csc.ReadContracts(contracsDir)

	sumMbPerSec := 0.00
	sumRatio := 0.00

	dict, err := ioutil.ReadFile(dictionary)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range contracts {
		start := time.Now()

		dst := make([]byte, 0)
		w := bytes.NewBuffer(dst)

		writer := zstd.NewWriterLevelDict(w, zstd.BestSpeed, dict)
		r := bytes.NewReader(c.Data)

		_, _ = io.Copy(writer, r)
		_ = writer.Close() // Make sure the writer is closed

		dst = w.Bytes()

		mbpersec := csc.CompressionSpeed(float64(len(c.Data)), start)
		sumMbPerSec = sumMbPerSec + mbpersec

		ratio := csc.CompressionRatio(float64(len(c.Data)), float64(len(dst)))
		sumRatio = sumRatio + ratio

		//t.Logf("Name: %s | UnCompressed Size: %d | Compressed Size: %d | Speed: %.2f MB/s | Ratio: %.2f", c.Name, c.Size, len(dst), mbpersec, ratio)
	}

	avgRatio := sumRatio / float64(len(contracts))
	avgSpeed := sumMbPerSec / float64(len(contracts))

	t.Logf("Average Compression Ratio: %.2f , Average Compression Speed: %.2f MB/s", avgRatio, avgSpeed)
}

func TestDecompressAllContracts(t *testing.T) {
	contracts := csc.ReadContracts(contracsDir)

	sumMbPerSec := 0.00
	for _, c := range contracts {
		compressed, err := zstd.Compress(nil, c.Data)
		if err != nil {
			t.Fatal(err)
		}

		start := time.Now()
		_, err = zstd.Decompress(nil, compressed)
		if err != nil {
			t.Fatal(err)
		}

		mbpersec := csc.CompressionSpeed(float64(len(c.Data)), start)
		sumMbPerSec = sumMbPerSec + mbpersec

		//t.Logf("Name: %s | Orig Size: %d | Compressed Size: %d | UnCompressed Size: %d | Speed: %.2f MB/s", c.Name, c.Size, len(compressed), len(dst), mbpersec)
	}

	avgSpeed := sumMbPerSec / float64(len(contracts))

	t.Logf("Average DeCompression Speed: %.2f MB/s", avgSpeed)
}

func TestNBATopShotContract(t *testing.T) {
	contracts := csc.ReadContracts(contracsDir)
	name := "0b2a3299cc857e29.TopShot.cdc"
	c := contracts[name]

	start := time.Now()

	dst := make([]byte, 0)
	dst, err := zstd.Compress(dst, c.Data)
	if err != nil {
		t.Fatal(err)
	}

	compressionTime := time.Since(start).Microseconds()
	compressionSpeed := csc.CompressionSpeed(float64(len(c.Data)), start)
	ratio := csc.CompressionRatio(float64(len(c.Data)), float64(len(dst)))

	start = time.Now()
	_, err = zstd.Decompress(nil, dst)
	if err != nil {
		t.Fatal(err)
	}

	decompressionTime := time.Since(start).Microseconds()
	decompressionSpeed := csc.CompressionSpeed(float64(len(c.Data)), start)

	t.Logf("compression speed: %.2f MB/s | compression time (micro seconds): %d ms | Orig Size: %d, Compression Size: %d | Compression Ratio: %.2f", compressionSpeed, compressionTime, len(c.Data), len(dst), ratio)
	t.Logf("decompression speed: %.2f MB/s | decompression time (micro seconds): %d ms", decompressionSpeed, decompressionTime)
}
