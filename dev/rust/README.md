Install
```
$ curl https://sh.rustup.rs -sSf | sh
```

Setup PATH
```bash
# ~/.config/fish/config.fish
set -x RUSTPATH $HOME/.cargo/
set -x PATH "$PATH:$RUSTPATH/bin"
```

Compile & Run
```
$ rustc sandbox.rs; ./sandbox
```
