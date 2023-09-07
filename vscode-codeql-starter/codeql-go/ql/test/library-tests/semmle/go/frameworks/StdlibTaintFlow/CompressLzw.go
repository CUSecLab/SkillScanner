// Code generated by https://github.com/gagliardetto/codebox. DO NOT EDIT.

package main

import (
	"compress/lzw"
	"io"
)

func TaintStepTest_CompressLzwNewReader_B0I0O0(sourceCQL interface{}) interface{} {
	fromReader656 := sourceCQL.(io.Reader)
	intoReadCloser414 := lzw.NewReader(fromReader656, 0, 0)
	return intoReadCloser414
}

func TaintStepTest_CompressLzwNewWriter_B0I0O0(sourceCQL interface{}) interface{} {
	fromWriteCloser518 := sourceCQL.(io.WriteCloser)
	var intoWriter650 io.Writer
	intermediateCQL := lzw.NewWriter(intoWriter650, 0, 0)
	link(fromWriteCloser518, intermediateCQL)
	return intoWriter650
}

func RunAllTaints_CompressLzw() {
	{
		source := newSource(0)
		out := TaintStepTest_CompressLzwNewReader_B0I0O0(source)
		sink(0, out)
	}
	{
		source := newSource(1)
		out := TaintStepTest_CompressLzwNewWriter_B0I0O0(source)
		sink(1, out)
	}
}
