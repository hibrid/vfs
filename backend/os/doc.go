/*
Package os built-in os lib VFS implementation.

Usage

Rely on github.com/hibrid/vfs/v6/backend

  import(
      "github.com/hibrid/vfs/v6/backend"
      "github.com/hibrid/vfs/v6/backend/os"
  )

  func UseFs() error {
      fs := backend.Backend(os.Scheme)
      ...
  }

Or call directly:

  import _os "github.com/hibrid/vfs/v6/backend/os"

  func DoSomething() {
      fs := &_os.FileSystem{}
      ...
  }

See Also

See: https://golang.org/pkg/os/
*/
package os
