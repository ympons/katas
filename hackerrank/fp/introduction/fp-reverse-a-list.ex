defmodule Solution do
    def input(read) do
        case :io.fread('', '~d') do
            {:ok, [x]} -> input([x | read])
            _eof -> read
        end
    end
    def main do
        input([]) |> Enum.each(&(IO.puts &1))
    end
end

Solution.main
