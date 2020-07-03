package codec

import "github.com/pingcap/ticdc/cdc/model"

var HEADER = []byte{'B', 'L', '0', '2'}

type RedEventBatchEncoder struct {
	canalEncoder EventBatchEncoder
}

func (r *RedEventBatchEncoder) AppendResolvedEvent(ts uint64) error {
	return r.canalEncoder.AppendResolvedEvent(ts)
}

func (r *RedEventBatchEncoder) AppendRowChangedEvent(e *model.RowChangedEvent) error {
	return r.canalEncoder.AppendRowChangedEvent(e)
}

func (r *RedEventBatchEncoder) AppendDDLEvent(e *model.DDLEvent) error {
	return r.canalEncoder.AppendDDLEvent(e)
}

func (r *RedEventBatchEncoder) Build() (key []byte, value []byte) {
	rawKey, rawValue := r.canalEncoder.Build()
	value = append([]byte{}, HEADER...) // deep copy
	value = append(value, rawValue...)
	return rawKey, value
}

func (r *RedEventBatchEncoder) Size() int {
	return len(HEADER) + r.canalEncoder.Size()
}

func NewRedEventBatchEncoder() EventBatchEncoder {
	c := NewCanalEventBatchEncoder()
	r := &RedEventBatchEncoder{
		canalEncoder: c,
	}
	return r
}
