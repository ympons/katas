defmodule Solution do
  def filter(n) do
    with {:ok, [x]} <- :io.fread('', '~d') do
        if x < n do
            :io.format('~p~n', [x])
        end
        filter(n)
    else
      _eof -> :ok
    end
  end
  def main do
    {:ok, [n]} = :io.fread('', '~d')
    filter(n)
  end
end

Solution.main
