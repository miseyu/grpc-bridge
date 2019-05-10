FROM envoyproxy/envoy-alpine:v1.10.0

RUN apk --no-cache add libintl && \
  apk --no-cache add --virtual .gettext gettext && \
  cp /usr/bin/envsubst /usr/local/bin/envsubst && \
  apk del .gettext

COPY ./config/envoy-sidecar.yaml /etc/envoy-sidecar.yaml
RUN mkdir /var/log/envoy/
CMD /usr/local/bin/envoy -c /etc/envoy-sidecar.yaml