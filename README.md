# unexpected_return_pc
runtime: unexpected return pc for runtime.sigpanic called

## fixed
- `go get -u -v golang.org/x/sys`

`upgraded golang.org/x/sys` => `v0.0.0-20210507161434-a76c4d0a0096`

```
commit a76c4d0a0096537dc565908b53073460d96c8539 (HEAD)
Author: Cherry Mui <cherryyz@google.com>
Date:   Thu May 6 21:57:46 2021 -0400

    unix: take address in assembly for Darwin syscall wrappers
    
    In Go 1.17 we will introduce a register-based ABI on some
    platforms, as well as ABI wrappers to bridge the ABIs. For Darwin
    syscall wrappers, it needs to be called directly, instead of
    through wrappers. Currently, it is written as that the syscall
    functions are defined in assembly and their addresses are taken
    from Go using funcPC. In Go 1.17 this will result in the address
    of the ABI wrapper, which is undesired.
    
    In the syscall package in the standard library we changed to use
    a compiler intrinsic internal/abi.FuncPCABI0 to take the address
    of the syscall function. But that is not available to this repo
    and not available in older versions of Go. Here we take a
    different approach: taking the address directly from assembly.
    This also ensures we get the address of the defined syscall
    function, not the ABI wrapper.
    
    Updates golang/go#45702.
    
    Change-Id: Ia7480d0fb0ca4fb9bf2f36d2deb1e3e5e4eb8284
    Reviewed-on: https://go-review.googlesource.com/c/sys/+/317894
    Trust: Cherry Mui <cherryyz@google.com>
    Run-TryBot: Cherry Mui <cherryyz@google.com>
    TryBot-Result: Go Bot <gobot@golang.org>
    Reviewed-by: Than McIntosh <thanm@google.com>
```
