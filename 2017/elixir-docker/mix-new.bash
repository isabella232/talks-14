$ mix phoenix.gen.presence

* creating web/channels/presence.ex

Add your new module to your supervision tree,
in lib/gastronokids.ex:

    children = [
      ...
      supervisor(Gastronokids.Presence, []),
    ]

You're all set! See the Phoenix.Presence docs for more details:
http://hexdocs.pm/phoenix/Phoenix.Presence.html

