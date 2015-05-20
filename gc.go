package gb

import "path/filepath"

// gc toolchain

type gcToolchain struct {
	goos, goarch         string
	gc, cc, ld, as, pack string
}

type gcoption struct {
	goos, goarch string
}

func (t *gcToolchain) Pack(afiles ...string) error {
	args := []string{"r"}
	args = append(args, afiles...)
	dir := filepath.Dir(afiles[0])
	return run(dir, t.pack, args...)
}
