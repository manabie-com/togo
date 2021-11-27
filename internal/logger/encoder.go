package logger

import (
"encoding/base64"
"encoding/json"
"fmt"
"math"
"sync"
"time"
"unicode/utf8"

"go.uber.org/zap/buffer"
"go.uber.org/zap/zapcore"
)

var _bufPool = buffer.NewPool()

var _consolePool = sync.Pool{New: func() interface{} {
	return &consoleEncoder{}
}}

var _sliceEncoderPool = sync.Pool{
	New: func() interface{} {
		return &sliceArrayEncoder{elems: make([]interface{}, 0, 2)}
	},
}

func getSliceEncoder() *sliceArrayEncoder {
	return _sliceEncoderPool.Get().(*sliceArrayEncoder)
}

func putSliceEncoder(e *sliceArrayEncoder) {
	e.elems = e.elems[:0]
	_sliceEncoderPool.Put(e)
}

func getConsoleEncoder() *consoleEncoder {
	return _consolePool.Get().(*consoleEncoder)
}

func putConsoleEncoder(enc *consoleEncoder) {
	enc.EncoderConfig = nil
	enc.buf = nil
	enc.spaced = false
	enc.openNamespaces = 0
	_consolePool.Put(enc)
}

type consoleEncoder struct {
	*zapcore.EncoderConfig
	buf            *buffer.Buffer
	spaced         bool
	openNamespaces int
}

// DefaultConsoleEncoder returns default console encoder
func DefaultConsoleEncoder() zapcore.Encoder {
	return NewConsoleEncoder(DefaultConsoleEncoderConfig)
}

// NewConsoleEncoder ...
func NewConsoleEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return newConsoleEncoder(cfg, false)
}

func newConsoleEncoder(cfg zapcore.EncoderConfig, spaced bool) *consoleEncoder {
	return &consoleEncoder{
		EncoderConfig: &cfg,
		buf:           _bufPool.Get(),
		spaced:        spaced,
	}
}

func (enc *consoleEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	enc.addKey(key)
	return enc.AppendArray(arr)
}

func (enc *consoleEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	enc.addKey(key)
	return enc.AppendObject(obj)
}

func (enc *consoleEncoder) AddBinary(key string, val []byte) {
	enc.AddString(key, base64.StdEncoding.EncodeToString(val))
}

func (enc *consoleEncoder) AddByteString(key string, val []byte) {
	enc.addKey(key)
	enc.AppendByteString(val)
}

func (enc *consoleEncoder) AddBool(key string, val bool) {
	enc.addKey(key)
	enc.AppendBool(val)
}

func (enc *consoleEncoder) AddComplex128(key string, val complex128) {
	enc.addKey(key)
	enc.AppendComplex128(val)
}

func (enc *consoleEncoder) AddDuration(key string, val time.Duration) {
	enc.addKey(key)
	enc.AppendDuration(val)
}

func (enc *consoleEncoder) AddFloat64(key string, val float64) {
	enc.addKey(key)
	enc.AppendFloat64(val)
}

func (enc *consoleEncoder) AddInt64(key string, val int64) {
	enc.addKey(key)
	enc.AppendInt64(val)
}

func (enc *consoleEncoder) AddReflected(key string, obj interface{}) error {
	marshaled, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	enc.addKey(key)
	_, err = enc.buf.Write(marshaled)
	return err
}

func (enc *consoleEncoder) OpenNamespace(key string) {
	enc.addKey(key)
	enc.buf.AppendByte('{')
	enc.openNamespaces++
}

func (enc *consoleEncoder) AddString(key, val string) {
	enc.addKey(key)
	enc.AppendString(val)
}

func (enc *consoleEncoder) AddTime(key string, val time.Time) {
	enc.addKey(key)
	enc.AppendTime(val)
}

func (enc *consoleEncoder) AddUint64(key string, val uint64) {
	enc.addKey(key)
	enc.AppendUint64(val)
}

func (enc *consoleEncoder) AppendArray(arr zapcore.ArrayMarshaler) error {
	enc.buf.AppendByte('[')
	err := arr.MarshalLogArray(enc)
	enc.buf.AppendByte(']')
	return err
}

func (enc *consoleEncoder) AppendObject(obj zapcore.ObjectMarshaler) error {
	enc.buf.AppendByte('{')
	err := obj.MarshalLogObject(enc)
	enc.buf.AppendByte('}')
	return err
}

func (enc *consoleEncoder) AppendBool(val bool) {
	enc.buf.AppendBool(val)
}

func (enc *consoleEncoder) AppendByteString(val []byte) {
	enc.buf.AppendByte('"')
	enc.safeAddByteString(val)
	enc.buf.AppendByte('"')
}

func (enc *consoleEncoder) AppendComplex128(val complex128) {
	r, i := float64(real(val)), float64(imag(val))
	enc.buf.AppendByte('"')
	enc.buf.AppendFloat(r, 64)
	enc.buf.AppendByte('+')
	enc.buf.AppendFloat(i, 64)
	enc.buf.AppendByte('i')
	enc.buf.AppendByte('"')
}

func (enc *consoleEncoder) AppendDuration(val time.Duration) {
	cur := enc.buf.Len()
	enc.EncodeDuration(val, enc)
	if cur == enc.buf.Len() {
		enc.AppendInt64(int64(val))
	}
}

func (enc *consoleEncoder) AppendInt64(val int64) {
	enc.buf.AppendInt(val)
}

func (enc *consoleEncoder) AppendReflected(val interface{}) error {
	marshaled, err := json.Marshal(val)
	if err != nil {
		return err
	}
	_, err = enc.buf.Write(marshaled)
	return err
}

func (enc *consoleEncoder) AppendString(val string) {
	enc.buf.AppendByte('"')
	enc.safeAddString(val)
	enc.buf.AppendByte('"')
}

func (enc *consoleEncoder) AppendTime(val time.Time) {
	cur := enc.buf.Len()
	enc.EncodeTime(val, enc)
	if cur == enc.buf.Len() {
		enc.AppendInt64(val.UnixNano())
	}
}

func (enc *consoleEncoder) AppendUint64(val uint64) {
	enc.buf.AppendUint(val)
}

func (enc *consoleEncoder) AddComplex64(k string, v complex64) { enc.AddComplex128(k, complex128(v)) }
func (enc *consoleEncoder) AddFloat32(k string, v float32)     { enc.AddFloat64(k, float64(v)) }
func (enc *consoleEncoder) AddInt(k string, v int)             { enc.AddInt64(k, int64(v)) }
func (enc *consoleEncoder) AddInt32(k string, v int32)         { enc.AddInt64(k, int64(v)) }
func (enc *consoleEncoder) AddInt16(k string, v int16)         { enc.AddInt64(k, int64(v)) }
func (enc *consoleEncoder) AddInt8(k string, v int8)           { enc.AddInt64(k, int64(v)) }
func (enc *consoleEncoder) AddUint(k string, v uint)           { enc.AddUint64(k, uint64(v)) }
func (enc *consoleEncoder) AddUint32(k string, v uint32)       { enc.AddUint64(k, uint64(v)) }
func (enc *consoleEncoder) AddUint16(k string, v uint16)       { enc.AddUint64(k, uint64(v)) }
func (enc *consoleEncoder) AddUint8(k string, v uint8)         { enc.AddUint64(k, uint64(v)) }
func (enc *consoleEncoder) AddUintptr(k string, v uintptr)     { enc.AddUint64(k, uint64(v)) }
func (enc *consoleEncoder) AppendComplex64(v complex64)        { enc.AppendComplex128(complex128(v)) }
func (enc *consoleEncoder) AppendFloat64(v float64)            { enc.appendFloat(v, 64) }
func (enc *consoleEncoder) AppendFloat32(v float32)            { enc.appendFloat(float64(v), 32) }
func (enc *consoleEncoder) AppendInt(v int)                    { enc.AppendInt64(int64(v)) }
func (enc *consoleEncoder) AppendInt32(v int32)                { enc.AppendInt64(int64(v)) }
func (enc *consoleEncoder) AppendInt16(v int16)                { enc.AppendInt64(int64(v)) }
func (enc *consoleEncoder) AppendInt8(v int8)                  { enc.AppendInt64(int64(v)) }
func (enc *consoleEncoder) AppendUint(v uint)                  { enc.AppendUint64(uint64(v)) }
func (enc *consoleEncoder) AppendUint32(v uint32)              { enc.AppendUint64(uint64(v)) }
func (enc *consoleEncoder) AppendUint16(v uint16)              { enc.AppendUint64(uint64(v)) }
func (enc *consoleEncoder) AppendUint8(v uint8)                { enc.AppendUint64(uint64(v)) }
func (enc *consoleEncoder) AppendUintptr(v uintptr)            { enc.AppendUint64(uint64(v)) }

func (enc *consoleEncoder) Clone() zapcore.Encoder {
	clone := enc.clone()
	clone.buf.Write(enc.buf.Bytes())
	return clone
}

func (enc *consoleEncoder) clone() *consoleEncoder {
	clone := getConsoleEncoder()
	clone.EncoderConfig = enc.EncoderConfig
	clone.spaced = enc.spaced
	clone.openNamespaces = enc.openNamespaces
	clone.buf = _bufPool.Get()
	return clone
}

func (enc *consoleEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line := _bufPool.Get()

	arr := getSliceEncoder()
	if enc.TimeKey != "" && enc.EncodeTime != nil {
		enc.EncodeTime(ent.Time, arr)
	}
	if enc.LevelKey != "" && enc.EncodeLevel != nil {
		enc.EncodeLevel(ent.Level, arr)
	}
	if enc.MessageKey != "" {
		enc.encodeMessage(ent.Message, arr)
	}
	if ent.Caller.Defined && enc.CallerKey != "" && enc.EncodeCaller != nil {
		enc.EncodeCaller(ent.Caller, arr)
	}
	for i := range arr.elems {
		if i > 0 {
			line.AppendByte('\t')
		}
		fmt.Fprint(line, arr.elems[i])
	}
	putSliceEncoder(arr)

	// Add any structured context.
	enc.writeContext(line, fields)

	// If there's no traceback key, honor that; this allows users to force single-line output.
	if ent.Stack != "" && enc.StacktraceKey != "" {
		line.AppendByte('\n')
		line.AppendString(ent.Stack)
	}

	if enc.LineEnding != "" {
		line.AppendString(enc.LineEnding)
	} else {
		line.AppendString(zapcore.DefaultLineEnding)
	}
	return line, nil
}

func (enc *consoleEncoder) encodeMessage(message string, e zapcore.PrimitiveArrayEncoder) {
	e.AppendString(message)
}

func (enc *consoleEncoder) writeContext(line *buffer.Buffer, extra []zapcore.Field) {
	context := enc.Clone().(*consoleEncoder)
	defer context.buf.Free()

	addFields(context, extra)
	context.closeOpenNamespaces()
	if context.buf.Len() == 0 {
		return
	}

	enc.addTabIfNecessary(line)
	line.Write(context.buf.Bytes())
}

func (enc *consoleEncoder) addTabIfNecessary(line *buffer.Buffer) {
	if line.Len() > 0 {
		line.AppendByte('\t')
	}
}

func (enc *consoleEncoder) truncate() {
	enc.buf.Reset()
}

func (enc *consoleEncoder) closeOpenNamespaces() {
	for i := 0; i < enc.openNamespaces; i++ {
		enc.buf.AppendByte('}')
	}
}

func (enc *consoleEncoder) addKey(key string) {
	enc.addElementSeparator()
	enc.safeAddString(key)
	enc.buf.AppendByte('=')
}

func (enc *consoleEncoder) addElementSeparator() {
	enc.buf.AppendByte('\t')
}

func (enc *consoleEncoder) appendFloat(val float64, bitSize int) {
	switch {
	case math.IsNaN(val):
		enc.buf.AppendString(`"NaN"`)
	case math.IsInf(val, 1):
		enc.buf.AppendString(`"+Inf"`)
	case math.IsInf(val, -1):
		enc.buf.AppendString(`"-Inf"`)
	default:
		enc.buf.AppendFloat(val, bitSize)
	}
}

func (enc *consoleEncoder) safeAddString(s string) {
	for i := 0; i < len(s); {
		if enc.tryAddRuneSelf(s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if enc.tryAddRuneError(r, size) {
			i++
			continue
		}
		enc.buf.AppendString(s[i : i+size])
		i += size
	}
}

func (enc *consoleEncoder) safeAddByteString(s []byte) {
	for i := 0; i < len(s); {
		if enc.tryAddRuneSelf(s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRune(s[i:])
		if enc.tryAddRuneError(r, size) {
			i++
			continue
		}
		enc.buf.Write(s[i : i+size])
		i += size
	}
}

func (enc *consoleEncoder) tryAddRuneSelf(b byte) bool {
	if b >= utf8.RuneSelf {
		return false
	}
	enc.buf.AppendByte(b)
	return true
}

func (enc *consoleEncoder) tryAddRuneError(r rune, size int) bool {
	if r == utf8.RuneError && size == 1 {
		enc.buf.AppendString(`\ufffd`)
		return true
	}
	return false
}

type sliceArrayEncoder struct {
	elems []interface{}
}

func (s *sliceArrayEncoder) AppendArray(v zapcore.ArrayMarshaler) error {
	enc := &sliceArrayEncoder{}
	err := v.MarshalLogArray(enc)
	s.elems = append(s.elems, enc.elems)
	return err
}

func (s *sliceArrayEncoder) AppendObject(v zapcore.ObjectMarshaler) error {
	m := zapcore.NewMapObjectEncoder()
	err := v.MarshalLogObject(m)
	s.elems = append(s.elems, m.Fields)
	return err
}

func (s *sliceArrayEncoder) AppendReflected(v interface{}) error {
	s.elems = append(s.elems, v)
	return nil
}

func (s *sliceArrayEncoder) AppendBool(v bool)              { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendByteString(v []byte)      { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendComplex128(v complex128)  { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendComplex64(v complex64)    { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendDuration(v time.Duration) { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendFloat64(v float64)        { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendFloat32(v float32)        { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendInt(v int)                { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendInt64(v int64)            { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendInt32(v int32)            { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendInt16(v int16)            { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendInt8(v int8)              { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendString(v string)          { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendTime(v time.Time)         { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendUint(v uint)              { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendUint64(v uint64)          { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendUint32(v uint32)          { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendUint16(v uint16)          { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendUint8(v uint8)            { s.elems = append(s.elems, v) }
func (s *sliceArrayEncoder) AppendUintptr(v uintptr)        { s.elems = append(s.elems, v) }

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}
