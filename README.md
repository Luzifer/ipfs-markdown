[![License: Apache v2.0](https://badge.luzifer.io/v1/badge?color=5d79b5&title=license&text=Apache+v2.0)](http://www.apache.org/licenses/LICENSE-2.0)
[![GoBuilder Download](https://badge.luzifer.io/v1/badge?color=5d79b5&title=Download&text=on GoBuilder)](https://gobuilder.me/github.com/Luzifer/ipfs-markdown)

# Luzifer / ipfs-markdown

This is a small Go webserver serving a static web interface (JavaScript) containing a Markdown parser and a code highlighter. The URLs are containing an [IPFS](http://ipfs.io/) hash which is fetched using a small proxy included in this webserver.

The intention behind this is to write a Markdown document and share an immutable version of it through IPFS. (For the background why these URLs are immutable please see the details about the IPFS project.)

## Usage

You need to be able to push documents into IPFS before using this viewer. After you uploaded your document (`ipfs add your_document.md`) you will get a hash of it. This hash is used in the URL of the viewer.

### Hosted version

A hosted version of this is available at https://markdown.luzifer.io/ and can be used with any hash containing a Markdown document. For example:

https://markdown.luzifer.io/QmPzC3XVh2umeYDqqkPEAeYjEEWMNC1ZfCMPsTNmZJweR4

### Your own copy

You can download a compiled version of the viewer on GoBuilder.me and just execute the binary. It will create a webserver on port 3000 and you can access your document there:

http://localhost:3000/QmPzC3XVh2umeYDqqkPEAeYjEEWMNC1ZfCMPsTNmZJweR4

This webserver can be deployed to any host in the internet and serve your Markdown files for you. If you want to you also can specify which IPFS gateway (public available IPFS http endpoint) to use for downloading the files:

```bash
# ./ipfs-markdown --help
Usage of ./ipfs-markdown:
      --branding="IPFS-Markdown-Viewer": Branding to show on page
      --gateway="http://gateway.ipfs.io": Gateway to fetch markdown from
      --listen=":3000": IP/Port to listen on
```
