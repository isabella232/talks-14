defmodule Kubernetes.Epmd do
  def start_link do
    :ignore
  end

  def register_node(name, port, _family) do
    register_node(name, port)
  end

  def register_node(_name, _port) do
    creation = :rand.uniform 3
    {:ok, creation}
  end

  def port_please(name, _ip) do
	port = Kubernetes.Proto.dist_port(name)
    version = 5
    {:port, port, version}
  end

  def names(_hostname) do
    {:error, :address}
  end
end
