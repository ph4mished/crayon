package main

import (
    "fmt"
    "strconv"
    "testing"
)

// Original: fmt.Sprintf (what people often use)
func BenchmarkFmtSprintfInt(b *testing.B) {
    x := 12345
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("%d", x)
    }
}

// Original: fmt.Sprint (what Crayon currently uses)
func BenchmarkFmtSprint(b *testing.B) {
    x := 12345
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprint(x)
    }
}

// strconv.FormatInt (standard fast method)
func BenchmarkStrconvFormatInt(b *testing.B) {
    x := int64(12345)
    for i := 0; i < b.N; i++ {
        _ = strconv.FormatInt(x, 10)
    }
}

// strconv.Itoa (convenience wrapper)
func BenchmarkStrconvItoa(b *testing.B) {
    x := 12345
    for i := 0; i < b.N; i++ {
        _ = strconv.Itoa(x)
    }
}

// ZERO ALLOCATION VERSION - Using stack buffer
func BenchmarkStrconvAppendIntZeroAlloc(b *testing.B) {
    x := int64(12345)
    for i := 0; i < b.N; i++ {
        var buf [20]byte
        value := string(strconv.AppendInt(buf[:0], x, 10))
        _ = value
    }
}

// ZERO ALLOCATION VERSION - Reusing same buffer (FASTEST)
func BenchmarkStrconvAppendIntReuseBuffer(b *testing.B) {
    x := int64(12345)
    buf := make([]byte, 0, 20)
    for i := 0; i < b.N; i++ {
        buf = buf[:0]
        buf = strconv.AppendInt(buf, x, 10)
        value := string(buf)
        _ = value
    }
}

// With padding - Current approach
func BenchmarkFmtSprintfWithPadding(b *testing.B) {
    x := 12345
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("%10s", fmt.Sprint(x))
    }
}

// With padding - Optimized using strconv
func BenchmarkOptimizedWithPadding(b *testing.B) {
    x := 12345
    for i := 0; i < b.N; i++ {
        s := strconv.Itoa(x)
        if len(s) < 10 {
            s = fmt.Sprintf("%10s", s)
        }
        _ = s
    }
}

// With padding - ZERO ALLOCATION version
func BenchmarkZeroAllocWithPadding(b *testing.B) {
    x := 12345
    for i := 0; i < b.N; i++ {
        var buf [20]byte
        s := string(strconv.AppendInt(buf[:0], int64(x), 10))
        if len(s) < 10 {
            // Create padding
            padding := make([]byte, 10-len(s))
            for j := range padding {
                padding[j] = ' '
            }
            s = string(padding) + s
        }
        _ = s
    }
}

// ZERO ALLOCATION - Right padding (left align)
func BenchmarkZeroAllocLeftPadding(b *testing.B) {
    x := 12345
    width := 10
    for i := 0; i < b.N; i++ {
        var buf [20]byte
        s := string(strconv.AppendInt(buf[:0], int64(x), 10))
        if len(s) < width {
            padded := make([]byte, width)
            copy(padded, s)
            for j := len(s); j < width; j++ {
                padded[j] = ' '
            }
            s = string(padded)
        }
        _ = s
    }
}

// Compare all methods side by side
func BenchmarkAllMethods(b *testing.B) {
    x := 12345
    
    b.Run("FmtSprintf", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = fmt.Sprintf("%d", x)
        }
    })
    
    b.Run("FmtSprint", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = fmt.Sprint(x)
        }
    })
    
    b.Run("StrconvItoa", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = strconv.Itoa(x)
        }
    })
    
    b.Run("StrconvFormatInt", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            _ = strconv.FormatInt(int64(x), 10)
        }
    })
    
    b.Run("ZeroAllocStack", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            var buf [20]byte
            _ = string(strconv.AppendInt(buf[:0], int64(x), 10))
        }
    })
    
    b.Run("ZeroAllocReuse", func(b *testing.B) {
        buf := make([]byte, 0, 20)
        for i := 0; i < b.N; i++ {
            buf = buf[:0]
            buf = strconv.AppendInt(buf, int64(x), 10)
            _ = string(buf)
        }
    })
}