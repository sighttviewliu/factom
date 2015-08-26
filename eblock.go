// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EBlock struct {
	Header struct {
		BlockSequenceNumber int
		ChainID             string
		PrevKeyMR           string
		Timestamp           uint64
	}
	EntryList []EBEntry
}

type EBEntry struct {
	Timestamp int64
	EntryHash string
}

func GetEBlock(keymr string) (*EBlock, error) {
	resp, err := http.Get(
		fmt.Sprintf("http://%s/v1/entry-block-by-keymr/%s", server, keymr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}

	e := new(EBlock)
	if err := json.Unmarshal(body, e); err != nil {
		return nil, err
	}

	return e, nil
}