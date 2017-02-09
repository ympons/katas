defmodule Solution do
  def pascal(n), do: Enum.reverse(pascal(n, [1,1], [[1]]))
  defp pascal(1, _, acc), do: acc
  defp pascal(n, row, acc) do
    next = Enum.map(Enum.zip(row ++ [0], [0 | row]), fn {a, b} -> a + b end)
    pascal(n-1, next, [row | acc])
  end
  def main do
    {:ok, [n]} = :io.fread('', '~d')
    x = pascal(n)
    Enum.each(x, &IO.puts "#{Enum.join(&1, " ")}")
  end
end

Solution.main
