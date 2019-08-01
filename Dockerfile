FROM scratch

COPY bin/soma /soma

ENTRYPOINT ["/soma"]


