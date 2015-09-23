FROM alpine

RUN apk --update add ca-certificates
ADD ./ipfs-markdown /ipfs-markdown

EXPOSE 3000
ENTRYPOINT ["/ipfs-markdown"]
CMD ["--"]
