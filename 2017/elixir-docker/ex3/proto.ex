defmodule Kubernetes.Proto do
    
      def dist_port(name) when is_atom(name) do
        dist_port Atom.to_string name
      end
    
      def dist_port(name) when is_list(name) do
        dist_port List.to_string name
      end
    
      def dist_port(name) when is_binary(name) do
        4370        
      end 
end
