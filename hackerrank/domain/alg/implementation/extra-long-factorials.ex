defmodule Solution do
    defp factorial(n, acc) do
        if n < 2, do: acc, else: factorial(n-1, n*acc)
    end
    def factorial(n), do: factorial(n, 1)
    def main do
        {:ok, [n]} = :io.fread('', '~d')
        IO.puts "#{factorial(n)}"
    end
end

Solution.main
