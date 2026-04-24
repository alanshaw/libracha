package capabilities

import (
	"encoding/json"
	"io"
	"time"

	jsg "github.com/alanshaw/dag-json-gen"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type CborTime time.Time

func (ct CborTime) MarshalCBOR(w io.Writer) error {
	nsecs := ct.Time().UnixNano()

	cbi := cbg.CborInt(nsecs)

	return cbi.MarshalCBOR(w)
}

func (ct *CborTime) UnmarshalCBOR(r io.Reader) error {
	var cbi cbg.CborInt
	if err := cbi.UnmarshalCBOR(r); err != nil {
		return err
	}

	t := time.Unix(0, int64(cbi))

	*ct = (CborTime)(t)
	return nil
}

func (ct CborTime) Time() time.Time {
	return (time.Time)(ct)
}

func (ct CborTime) MarshalDagJSON(w io.Writer) error {
	nsecs := ct.Time().UnixNano()
	buf, err := json.Marshal(nsecs)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}

func (ct *CborTime) UnmarshalDagJSON(r io.Reader) error {
	var nsecs int64
	jr := jsg.NewDagJsonReader(r)
	nsecs, err := jr.ReadNumberAsInt64()
	if err != nil {
		return err
	}
	t := time.Unix(0, nsecs)
	*ct = (CborTime)(t)
	return nil
}

func (ct CborTime) MarshalJSON() ([]byte, error) {
	return ct.Time().MarshalJSON()
}

func (ct *CborTime) UnmarshalJSON(b []byte) error {
	var t time.Time
	if err := t.UnmarshalJSON(b); err != nil {
		return err
	}
	*(*time.Time)(ct) = t
	return nil
}
