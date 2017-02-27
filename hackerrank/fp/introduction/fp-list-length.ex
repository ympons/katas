defmodule Solution do
    def input do
        case :io.fread('','~d') do
            {:ok, [x]} -> [x | input()]
            _eof -> []
        end
    end
    defp count([], count), do: count
    defp count([_head | tail], count), do: count(tail, count+1)
    def count(l), do: count(l, 0)
    def main do
        IO.puts input |> count
    end
end

Solution.main
