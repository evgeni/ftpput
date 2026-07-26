# ftpput - minimal ftp server that allows uploads and nothing else

## why?

I needed a way for my Brother scanner to drop files into a folder and I wanted as few other features as possible.

## configuration

Set the `FTPPUT_DIR` environment variable to store incoming files somewhere else.
If unset, it defaults to `.`, so *probably* not what you want.
