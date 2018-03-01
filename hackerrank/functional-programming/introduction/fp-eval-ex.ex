defmodule Solution do
    defp f(n, acc) do
        if n < 2, do: acc, else: f(n-1, n*acc)
    end
    defp f(n), do: f(n, 1)

    defp pow(x, i, acc) do
        if i == 0, do: acc, else: pow(x, i-1, x*acc)
    end
    def pow(x, i), do: pow(x, i, 1)
    
    def e(n) do
        Enum.map(1..9, &(pow(n, &1)/f(&1)))
            |> Enum.reduce(0, &(&1 + &2))
    end
    
    def solve(t) do
        if t > 0 do
            {:ok, [x]} = :io.fread('', '~f')
            IO.puts 1 + e(x)
            solve(t-1)
        end
    end
    
    def main do
        {:ok, [t]} = :io.fread('','~d')
        solve(t)
    end
end

Solution.main
