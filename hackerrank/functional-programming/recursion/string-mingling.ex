defmodule Solution do
    def main do
        a = :io.get_line('') |> String.codepoints
        b = :io.get_line('') |> String.codepoints
        Enum.zip(a,b) 
          |> Enum.flat_map(&[elem(&1,0),elem(&1,1)])
          |> IO.puts
    end
end

Solution.main
