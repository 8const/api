// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "0cf9b1df4a6e0b0259c81be8715c8be1"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"e18474c22fdb011182a2afd17d522f09": "1f8b08000000000000ff8ccd410e82301085e1fd9ce2edd0989e80150a3b1215e100482732b1b6135a83c737c46874e7f6e5fdf98cc1e62697a94f8c4ec9189c8e35c423f2902478649d6690087ef0704f6c318fec9146897875cb49227a55276c69d754455ba12db67505e5a08eb1120bf1699d13d1375886d9bfc98fb78c7f8953708e2dcefd70a5b2d91f7ecc9c9e010000ffff0d705ef9da000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("migrations", packr.Pointer{ForwardBox: gk, ForwardPath: "e18474c22fdb011182a2afd17d522f09"})
	}()
	return nil
}()
