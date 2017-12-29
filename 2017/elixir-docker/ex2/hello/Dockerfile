FROM elixir:1.5 AS build
COPY . /build
WORKDIR /build
ENV MIX_ENV=prod
RUN mix local.hex --force
RUN mix local.rebar --force
RUN mix deps.get
RUN mix phx.digest
RUN mix release --env=prod

FROM debian:jessie-slim
WORKDIR /run
RUN DEBIAN_FRONTEND=noninteractive apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y openssl
COPY --from=build /build/_build/prod/rel/hello .
ENV LANG C.UTF-8
ENV REPLACE_OS_VARS=true
ENV PORT 4000
ENV MIX_ENV=prod
ENTRYPOINT ["/run/bin/hello", "foreground"]
EXPOSE 4000
