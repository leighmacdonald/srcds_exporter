FROM chromedp/headless-shell:latest
LABEL maintainer="Alexander Trost <galexrt@googlemail.com>"
WORKDIR /
ADD srcds_exporter /bin/srcds_exporter
ENTRYPOINT ["/bin/srcds_exporter"]
CMD ["-collectors.enabled", "map,players,rank"]
