# srn.go
Create new file with replaced strings

## Installation

```bash
> git clone https://github.com/rahilwazir/srn.go.git
> cd srn.go
> go build srn.go
> sudo ln -s srn /usr/local/bin/
```

## Usage

Replace keyword `google` with `facebook`

```bash
> srn /tmp/google.conf google facebook
New file created at: tmp/facebook.conf
```

Or you can skip second parameter by adding `_`, which indicates the basename `google` of file `google.conf`

```bash
> srn /tmp/google.conf _ facebook
New file created at: tmp/facebook.conf
```

Practicing golang, plus I wanted a tool which can do something similar to this (I know `sed` but still mmm) because I have to repeat and replace text for my nginx vhost config, now I create vhost easily :D
