# srn.go
Create new file with replaced strings

## Installation

```
git clone https://github.com/rahilwazir/srn.go.git
cd srn.go
go build srn.go
sudo ln -s srn /usr/local/bin/
```

## Usage

```bash
$ srn /tmp/google.conf google facebook // replace google with facebook
New file created at: tmp/facebook.conf

// or
$ srn /tmp/google.conf _ facebook // _ indicates basename of file "google" (wihtout ext)
New file created at: tmp/facebook.conf
```

Practicing golang, plus I wanted a tool which can do something similar to this (I know `sed` but still mmm) because I have to repeat and replace text for my nginx vhost config, now I create vhost easily :D