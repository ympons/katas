defmodule Solution do
  def gcd(a, 0), do: a
  def gcd(a, b), do: gcd(b, rem(a,b))
  def read(n) do
    Enum.reduce(1..n, 1, fn(_, acc)->
      {:ok, [v]} = :io.fread('', '~d')
      acc * v
    end)
  end
  def main do
    {:ok, [n]} = :io.fread('', '~d')
    a = read(n)
    {:ok, [m]} = :io.fread('', '~d')
    b = read(m)
    gcd(a, b) |> rem(1000000007) |> IO.puts
  end
end

Solution.main
