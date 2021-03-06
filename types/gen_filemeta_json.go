// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/anacrolix/torrent/metainfo"
)

// MarshalJSON marshals as JSON.
func (f FileMeta) MarshalJSON() ([]byte, error) {
	type FileMeta struct {
		InfoHash metainfo.Hash `json:"infoHash"         gencodec:"required"`
		RawSize  uint64        `json:"rawSize"          gencodec:"required"`
	}
	var enc FileMeta
	enc.InfoHash = f.InfoHash
	enc.RawSize = f.RawSize
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *FileMeta) UnmarshalJSON(input []byte) error {
	type FileMeta struct {
		InfoHash *metainfo.Hash `json:"infoHash"         gencodec:"required"`
		RawSize  *uint64        `json:"rawSize"          gencodec:"required"`
	}
	var dec FileMeta
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.InfoHash == nil {
		return errors.New("missing required field 'infoHash' for FileMeta")
	}
	f.InfoHash = *dec.InfoHash
	if dec.RawSize == nil {
		return errors.New("missing required field 'rawSize' for FileMeta")
	}
	f.RawSize = *dec.RawSize
	return nil
}
