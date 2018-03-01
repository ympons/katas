defmodule Solution do 
    def main do
        {:ok, [n]} = :io.fread('','~d')
        l = Enum.map(1..n, &(&1))
        :io.format('~w~n', [l])
    end
end

Solution.main
