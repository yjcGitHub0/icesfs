package vfs

import (
	"context"
	"github.com/go-playground/assert/v2"
	"icesos/command/vars"
	"icesos/entry"
	"icesos/full_path"
	"icesos/kv"
	"icesos/kv/redis_store"
	"icesos/set"
	"icesos/storage_engine"
	"icesos/util"
	"os"
	"testing"
	"time"
)

func TestEntry(t *testing.T) {
	kvStore := redis_store.NewRedisStore(vars.RedisHostPost, vars.RedisPassword, vars.RedisDatabase)
	storageEngine := storage_engine.NewStorageEngine(vars.MasterServer)
	vfs := NewVFS(kvStore, storageEngine)

	fp := full_path.FullPath("/aa/bb/cc")
	setName := set.Set("test")
	ctx := context.Background()

	size := uint64(5 * 1024)

	fid := putObject(t, ctx, vfs, size)

	ent := &entry.Entry{
		FullPath: fp,
		Set:      setName,
		Ctime:    time.Unix(time.Now().Unix(), 0), // windows: precision to s
		Mode:     os.ModePerm,
		Mime:     "",
		Md5:      util.RandMd5(),
		FileSize: size,
		Fid:      fid,
	}

	assert.Equal(t, ent.IsFile(), true)
	assert.Equal(t, ent.IsDirectory(), false)

	err := vfs.insertEntry(ctx, ent)
	assert.Equal(t, err, nil)

	ent2, err := vfs.getEntry(ctx, setName, fp)
	assert.Equal(t, err, nil)
	assert.Equal(t, ent2, ent)
	assert.Equal(t, ent2.IsFile(), true)
	assert.Equal(t, ent2.IsDirectory(), false)

	err = vfs.deleteEntry(ctx, setName, fp)
	assert.Equal(t, err, nil)

	entry3, err := vfs.getEntry(ctx, setName, fp)
	assert.Equal(t, err, kv.KvNotFound)
	assert.Equal(t, entry3, nil)

	err = vfs.deleteEntry(ctx, setName, fp)
	assert.Equal(t, err, nil)
}