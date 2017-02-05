defmodule Solution do
    def input do
        case :io.fread('', '~d') do
            {:ok, [x]} -> [x | input()]
            _eof -> []
        end
    end
    def main do
        input 
            |> Enum.filter(&(rem(&1, 2) != 0))
            |> Enum.reduce(0, &(&1 + &2))
            |> IO.puts
    end
end

Solution.main
