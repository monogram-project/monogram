package lib

import "fmt"

type LineCol struct {
	LineNo int // The starting line number of the token
	ColNo  int // The starting column number of the token
}

type Span struct {
	StartLine   int // The starting line number of the token
	StartColumn int // The starting column number of the token
	EndLine     int // The ending line number of the token
	EndColumn   int // The ending column number of the token
}

func (x *Span) SpanString() string {
	return fmt.Sprintf("%d %d %d %d", x.StartLine, x.StartColumn, x.EndLine, x.EndColumn)
}

func (x *LineCol) SpanString(lineCol LineCol) string {
	span := Span{
		StartLine:   x.LineNo,
		StartColumn: x.ColNo,
		EndLine:     lineCol.LineNo,
		EndColumn:   lineCol.ColNo,
	}
	return span.SpanString()
}

func (x *LineCol) Span(lineCol LineCol) Span {
	return Span{
		StartLine:   x.LineNo,
		StartColumn: x.ColNo,
		EndLine:     lineCol.LineNo,
		EndColumn:   lineCol.ColNo,
	}
}
