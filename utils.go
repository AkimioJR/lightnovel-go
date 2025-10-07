package lightnovel

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

// 支持gzip -> base64 -> zlib的解压链
func decompressResponse(resp *http.Response) ([]byte, error) {
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("create gzip reader失败: %w", err)
	}
	defer gzipReader.Close()

	// 读取第一层解压后的数据
	data, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("read gzip data failed: %w", err)
	}

	// 第二层：检查并处理base64编码的数据
	decodedData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, fmt.Errorf("decode base64 failed: %w", err)
	}

	// 第三层：处理zlib压缩
	zlibReader, err := zlib.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return nil, fmt.Errorf("create zlib reader failed: %w", err)
	}
	defer zlibReader.Close()

	finalData, err := io.ReadAll(zlibReader)
	if err != nil {
		return nil, fmt.Errorf("decompress zlib failed: %w", err)
	}
	return finalData, nil
}
