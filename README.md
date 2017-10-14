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

```
srn /tmp/file.conf google.dev facebook.dev // replace google.dev with facebook.dev
// new file created at /tmp/facebook.conf on success
```

Practicing golang, plus I wanted a tool which can do something similar to this (I know `sed` but still mmm) because I have to repeat and replace text for my nginx vhost config, now I create vhost easily :D