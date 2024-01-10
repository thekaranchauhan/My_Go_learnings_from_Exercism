package paasio

import (
	"io"
	"sync"
)

type counter struct {
	mu        sync.RWMutex
	byteCount int64
	opCount   int
}
type rwCounter struct {
	r            io.Reader
	w            io.Writer
	rCntr, wCntr counter
}

func (c *counter) count() (int64, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.byteCount, c.opCount
}

func (c *counter) update(byteCount int, err error) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.byteCount += int64(byteCount)
	c.opCount++
	return byteCount, err
}

func NewReadCounter(r io.Reader) ReadCounter                { return &rwCounter{r: r} }
func NewWriteCounter(w io.Writer) WriteCounter              { return &rwCounter{w: w} }
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter { return &rwCounter{r: rw, w: rw} }

func (rwc *rwCounter) Read(p []byte) (int, error)  { return rwc.rCntr.update(rwc.r.Read(p)) }
func (rwc *rwCounter) ReadCount() (int64, int)     { return rwc.rCntr.count() }
func (rwc *rwCounter) Write(p []byte) (int, error) { return rwc.wCntr.update(rwc.w.Write(p)) }
func (rwc *rwCounter) WriteCount() (int64, int)    { return rwc.wCntr.count() }
