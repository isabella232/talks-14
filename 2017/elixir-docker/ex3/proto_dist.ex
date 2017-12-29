defmodule Kubernetes.Proto_dist do
  def listen(name) do
    port = Kubernetes.Proto.dist_port(name)
    # Set both "min" and "max" variables, to force the port number to this one.
    :ok = :application.set_env :kernel, :inet_dist_listen_min, port
    :ok = :application.set_env :kernel, :inet_dist_listen_max, port
    :inet_tcp_dist.listen name
  end

  def select(node), do: :inet_tcp_dist.select node
  def accept(listen), do: :inet_tcp_dist.accept listen

  def accept_connection(accept_pid, socket, my_node, allowed, setup_time) do
    :inet_tcp_dist.accept_connection accept_pid, socket, my_node, allowed, setup_time
  end

  def setup(node, type, my_node, long_or_short_names, setup_time) do
    :inet_tcp_dist.setup node, type, my_node, long_or_short_names, setup_time
  end

  def close(listen), do: :inet_tcp_dist.close listen

  def childspecs, do: []
end