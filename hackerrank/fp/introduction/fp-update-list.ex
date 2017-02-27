defmodule Solution do
    def input do
        case :io.fread('', '~d') do
            {:ok, [x]} -> [x | input]
            _eof -> []
        end
    end
    def main do
        input 
            |> Enum.map(&(if &1 < 0, do: -1*&1, else: &1))
            |> Enum.each(&(IO.puts &1))
    end
end

Solution.main
