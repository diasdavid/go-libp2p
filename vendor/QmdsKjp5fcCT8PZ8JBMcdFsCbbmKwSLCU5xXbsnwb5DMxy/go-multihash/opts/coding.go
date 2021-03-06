package opts

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	base58 "QmNsoHoCVhgXcv1Yg45jtkMgimxorTAN36fV9AQMFXHHAQ/go-base58"
	mh "QmdsKjp5fcCT8PZ8JBMcdFsCbbmKwSLCU5xXbsnwb5DMxy/go-multihash"
)

func Decode(encoding, digest string) (mh.Multihash, error) {
	switch encoding {
	case "raw":
		return mh.Cast([]byte(digest))
	case "hex":
		return hex.DecodeString(digest)
	case "base58":
		return base58.Decode(digest), nil
	case "base64":
		return base64.StdEncoding.DecodeString(digest)
	default:
		return nil, fmt.Errorf("unknown encoding: %s", encoding)
	}
}

func Encode(encoding string, hash mh.Multihash) (string, error) {
	switch encoding {
	case "raw":
		return string(hash), nil
	case "hex":
		return hex.EncodeToString(hash), nil
	case "base58":
		return base58.Encode(hash), nil
	case "base64":
		return base64.StdEncoding.EncodeToString(hash), nil
	default:
		return "", fmt.Errorf("unknown encoding: %s", encoding)
	}
}
