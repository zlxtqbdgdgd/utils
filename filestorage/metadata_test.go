// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package filestorage_test

import (
	"time"

	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/utils/filestorage"
)

var (
	_ filestorage.Document = (*filestorage.DocWrapper)(nil)
	_ filestorage.Metadata = (*filestorage.FileMetadata)(nil)
)

var _ = gc.Suite(&MetadataSuite{})

type MetadataSuite struct {
	testing.IsolationSuite
}

func (s *MetadataSuite) TestFileMetadataNewMetadata(c *gc.C) {
	meta := filestorage.NewMetadata()

	c.Check(meta.ID(), gc.Equals, "")
	c.Check(meta.Size(), gc.Equals, int64(0))
	c.Check(meta.Checksum(), gc.Equals, "")
	c.Check(meta.ChecksumFormat(), gc.Equals, "")
	c.Check(meta.Stored(), gc.IsNil)
}

func (s *MetadataSuite) TestFileMetadataDoc(c *gc.C) {
	meta := filestorage.NewMetadata()
	meta.SetFile(10, "some sum", "SHA-1")
	doc := meta.Doc()

	c.Check(doc, gc.Equals, meta)
}

func (s *MetadataSuite) TestFileMetadataSetIDInitial(c *gc.C) {
	meta := filestorage.NewMetadata()
	meta.SetFile(10, "some sum", "SHA-1")
	c.Assert(meta.ID(), gc.Equals, "")

	success := meta.SetID("some id")
	c.Check(success, gc.Equals, false)
	c.Check(meta.ID(), gc.Equals, "some id")
}

func (s *MetadataSuite) TestFileMetadataSetIDAlreadySetSame(c *gc.C) {
	meta := filestorage.NewMetadata()
	meta.SetFile(10, "some sum", "SHA-1")
	success := meta.SetID("some id")
	c.Assert(success, gc.Equals, false)

	success = meta.SetID("some id")
	c.Check(success, gc.Equals, true)
	c.Check(meta.ID(), gc.Equals, "some id")
}

func (s *MetadataSuite) TestFileMetadataSetIDAlreadySetDifferent(c *gc.C) {
	meta := filestorage.NewMetadata()
	meta.SetFile(10, "some sum", "SHA-1")
	success := meta.SetID("some id")
	c.Assert(success, gc.Equals, false)

	success = meta.SetID("another id")
	c.Check(success, gc.Equals, true)
	c.Check(meta.ID(), gc.Equals, "some id")
}

func (s *MetadataSuite) TestFileMetadataSetFile(c *gc.C) {
	meta := filestorage.NewMetadata()
	c.Assert(meta.Size(), gc.Equals, int64(0))
	c.Assert(meta.Checksum(), gc.Equals, "")
	c.Assert(meta.ChecksumFormat(), gc.Equals, "")
	c.Assert(meta.Stored(), gc.IsNil)
	meta.SetFile(10, "some sum", "SHA-1")

	c.Check(meta.Size(), gc.Equals, int64(10))
	c.Check(meta.Checksum(), gc.Equals, "some sum")
	c.Check(meta.ChecksumFormat(), gc.Equals, "SHA-1")
	c.Check(meta.Stored(), gc.IsNil)
}

func (s *MetadataSuite) TestFileMetadataSetStored(c *gc.C) {
	meta := filestorage.NewMetadata()
	timestamp := time.Now().UTC()
	meta.SetStored(&timestamp)

	c.Check(meta.Stored(), gc.Equals, &timestamp)
}

func (s *MetadataSuite) TestFileMetadataSetStoredDefault(c *gc.C) {
	meta := filestorage.NewMetadata()
	c.Assert(meta.Stored(), gc.IsNil)
	meta.SetStored(nil)

	c.Check(meta.Stored(), gc.NotNil)
}

func (s *MetadataSuite) TestFileMetadataCopy(c *gc.C) {
	meta := filestorage.NewMetadata()
	meta.SetFile(10, "some sum", "SHA-1")
	doc := meta.Copy("")
	copied, ok := doc.(filestorage.Metadata)
	c.Assert(ok, jc.IsTrue)

	c.Check(copied, gc.Not(gc.Equals), meta)
	c.Check(copied, gc.DeepEquals, meta)
}
