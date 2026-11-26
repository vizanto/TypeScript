package main

import "github.com/microsoft/TypeScript/tsc/internal/batch"

// batchSys is the CLI system; it opts batch compilations into process-wide runtime tuning.
type batchSys struct {
	*osSys
	programBound func()
}

func (s *batchSys) BeforeBatchCompilation() func() {
	programBound, end := batch.Begin()
	s.programBound = programBound
	return end
}

func (s *batchSys) ProgramBound() {
	if s.programBound != nil {
		s.programBound()
	}
}
