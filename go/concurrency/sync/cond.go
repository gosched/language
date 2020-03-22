package explore

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

// MyBucket .
type MyBucket struct {
	buffer *bytes.Buffer
	mutex  *sync.RWMutex
	cond   *sync.Cond
}

// NewBucket .
func NewBucket() *MyBucket {
	buf := make([]byte, 0)
	bucket := &MyBucket{
		buffer: bytes.NewBuffer(buf),
		mutex:  new(sync.RWMutex),
	}
	bucket.cond = sync.NewCond(bucket.mutex.RLocker())
	return bucket
}

func (bucket *MyBucket) Read(i int) {
	bucket.mutex.RLock()
	defer bucket.mutex.RUnlock()
	var data []byte
	var d byte
	var err error
	for {
		if d, err = bucket.buffer.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				bucket.cond.Wait()
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
	}
}

// Write .
func (bucket *MyBucket) Write(d []byte) (int, error) {
	bucket.mutex.Lock()
	defer bucket.mutex.Unlock()
	n, err := bucket.buffer.Write(d)
	bucket.cond.Broadcast()
	return n, err
}

// HelloBucket .
func HelloBucket() {
	bucket := NewBucket()

	go bucket.Read(1)
	go bucket.Read(2)

	for i := 0; i < 10; i++ {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			bucket.Write([]byte(d))
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
}
