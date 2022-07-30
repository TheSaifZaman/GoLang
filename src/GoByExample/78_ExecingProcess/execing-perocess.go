// In the previous example we looked at
// [spawning external processes](spawning-processes). We
// do this when we need an external process accessible to
// a running Go process. Sometimes we just want to
// completely replace the current Go process with another
// (perhaps non-Go) one. To do this we'll use Go's
// implementation of the classic
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// function.

package execingProcesses

import (
	"os"
	"os/exec"
	"syscall"
)

func ExecingProcesses() {

	// For our example we'll exec `ls`. Go requires an
	// absolute path to the binary we want to execute, so
	// we'll use `exec.LookPath` to find it (probably
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` requires arguments in slice form (as
	// opposed to one big string). We'll give `ls` a few
	// common arguments. Note that the first argument should
	// be the program name.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	env := os.Environ()

	// Here's the actual `syscall.Exec` call. If this call is
	// successful, the execution of our process will end
	// here and be replaced by the `/bin/ls -a -l -h`
	// process. If there is an error we'll get a return
	// value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
/*
total 45K    
drwxr-xr-x    2 root     root        4.0K Jul 30 20:27 .
drwxr-xr-x    2 root     root        4.0K Jul 30 20:27 ..
-rwxr-xr-x    1 root     root           0 Jul 30 20:27 .dockerenv
drwxr-xr-x    2 root     root       12.0K Feb 18  2021 bin
drwxr-xr-x    5 root     root         320 Jul 30 20:27 dev
drwxr-xr-x    1 root     root        4.0K Jul 30 20:27 etc
drwxr-xr-x    2 nobody   nobody      4.0K Mar  9  2021 home
drwxr-xr-x    2 root     root        4.0K Feb 18  2021 lib
lrwxrwxrwx    1 root     root           3 Feb 18  2021 lib64 -> lib
dr-xr-xr-x    5 root     root           0 Jul 30 20:27 proc
drwx------    2 root     root        4.0K Mar  9  2021 root
drwxr-xr-x   12 root     root           0 Jul 30 20:27 sys
drwxrwxrwt    2 root     root          40 Jul 30 20:27 tmp
drwxrwxrwt    2 root     root          60 Jul 30 20:28 tmpfs
drwxr-xr-x    1 root     root        4.0K Jul 23  2021 usr
drwxr-xr-x    4 root     root        4.0K Mar  9  2021 var
*/
