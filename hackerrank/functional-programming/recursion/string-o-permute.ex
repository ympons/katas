defmodule Solution do
    def permute([], b), do: Enum.reverse(b)
    def permute([x, y | tail], b), do: permute(tail, [x, y | b])
    def permute(a), do: permute(a, [])

    def solve(n) do
        if n > 0 do
            {:ok, [s]} = :io.fread('', '~s')
            s |> to_string
              |> String.codepoints
              |> permute
              |> IO.puts
            solve(n-1)
        end
    end
    def main do
        {:ok, [n]} = :io.fread('', '~d')
        solve(n)
    end
end

Solution.main
