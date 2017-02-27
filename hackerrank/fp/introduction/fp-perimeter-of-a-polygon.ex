defmodule Solution do
    def d({x1, y1}, {x2, y2}) do
        :math.sqrt(:math.pow(x1-x2, 2) + :math.pow(y1-y2, 2))
    end
    def input(0, ps), do: ps
    def input(n, ps) do
        {:ok, [x, y]} = :io.fread('', '~d~d')
        input(n-1, [{x, y} | ps])
    end
    def main do
        {:ok, [n]} = :io.fread('', '~d')
        points = input(n, [])
        [head | _] = points
        
        {_, perimeter} = points
            |> Enum.reverse
            |> Enum.reduce({head, 0}, fn(p, acc) -> {p, elem(acc, 1) + d(p, elem(acc,0))} end )
        IO.puts perimeter
    end
end

Solution.main
