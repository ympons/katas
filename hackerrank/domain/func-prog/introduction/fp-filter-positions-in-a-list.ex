defmodule Solution do
    defp read do
        IO.read(:all)
            |> String.trim
            |> String.split("\n")
            |> Enum.map(&String.to_integer/1)
    end
    
    defp filter(_, 0, x), do: x
    defp filter([head | tail], n, list) do
        if rem(n, 2) != 0, 
            do: filter(tail, n-1, [head | list]),
            else: filter(tail, n-1, list)
    end
    
    def main do
        list = read
        c = Enum.count(list)
        s = if rem(c, 2) == 0, do: c, else: c-1
        list |> filter(s, [])
             |> Enum.reverse
             |> Enum.each(&(IO.puts &1))
    end
end

Solution.main
