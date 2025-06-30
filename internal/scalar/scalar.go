package scalar

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalDate serializes a time.Time to a YYYY-MM-DD string literal.
func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		// write the date as a JSON string literal in YYYY-MM-DD
		fmt.Fprintf(w, "%q", t.Format("2006-01-02"))
	})
}

// UnmarshalDate parses a “YYYY-MM-DD” string into time.Time.
func UnmarshalDate(v interface{}) (time.Time, error) {
	s, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("date must be a string")
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid Date %q: expected format YYYY-MM-DD", s)
	}
	// ensure it round-trips exactly
	if t.Format("2006-01-02") != s {
		return time.Time{}, fmt.Errorf("invalid Date %q: expected format YYYY-MM-DD", s)
	}
	return t, nil
}

func ConvertUintPtrToIntPtr[T ~uint8 | ~uint16 | ~uint32](p *T) *int {
	if p == nil {
		return nil
	}
	v := int(*p)
	return &v
}

func ConvertInt8PtrToIntPtr(p *int8) *int {
	if p == nil {
		return nil
	}
	v := int(*p)
	return &v
}

func ConvertSliceToIntPtrSlice[T ~uint8 | ~uint16 | ~uint32](in []T) []*int {
	out := make([]*int, len(in))
	for i, v := range in {
		iv := int(v)
		out[i] = &iv
	}
	return out
}

func ConvertUint32PtrSliceToIntPtrSlice(in []*uint32) []*int {
	out := make([]*int, len(in))
	for i, p := range in {
		if p == nil {
			// preserve the null
			out[i] = nil
		} else {
			v := int(*p)
			out[i] = &v
		}
	}
	return out
}

// ConvertUint16PtrSliceToIntPtrSlice turns a []*uint16 into []*int
func ConvertUint16PtrSliceToIntPtrSlice(in []*uint16) []*int {
	out := make([]*int, len(in))
	for i, p := range in {
		if p == nil {
			out[i] = nil
		} else {
			v := int(*p)
			out[i] = &v
		}
	}
	return out
}

func ConvertFloat32PtrToFloat64Ptr(p *float32) *float64 {
	if p == nil {
		return nil
	}
	v := float64(*p)
	return &v
}

func ConvertFloat32PtrSliceToFloat64PtrSlice(in []*float32) []*float64 {
	out := make([]*float64, len(in))
	for i, p := range in {
		if p == nil {
			out[i] = nil
		} else {
			v := float64(*p)
			out[i] = &v
		}
	}
	return out
}
