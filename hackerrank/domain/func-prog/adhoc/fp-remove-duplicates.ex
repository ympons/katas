defmodule Solution do
    def main do
        :io.get_line('')
            |> String.strip
            |> String.codepoints
            |> Enum.uniq
            |> IO.puts
    end
end

Solution.main
