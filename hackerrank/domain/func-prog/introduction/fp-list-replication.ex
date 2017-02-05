defmodule Solution do
  def duplicate(n) do
    with {:ok, [x]} <- :io.fread('', '~d') do
        Enum.each(1..n, fn _ -> :io.format('~p~n', [x]) end)
        duplicate(n)
    else
      _eof -> :ok
    end
  end
  def main do
    {:ok, [n]} = :io.fread('', '~d')
    duplicate(n)
  end
end

Solution.main
