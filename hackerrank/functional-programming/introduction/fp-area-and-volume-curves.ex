defmodule Solution do
    def read do
        :io.get_line('')
            |> String.strip 
            |> String.split(" ") 
            |> Enum.map(&String.to_integer(&1))
    end
    
    def f(p, x) do
        p |> Enum.reduce(0, fn(i, acc)-> acc + elem(i,0)*(:math.pow(x,elem(i,1))) end)
    end
    
    def area(p, x) do
        0.001 * :math.pi * :math.pow(f(p,x), 2)
    end
    
    def sum(p, xi, func) do
        xi |> Enum.reduce(0, fn(x, acc)-> acc + func.(p, x) end)
    end
    
    def main do
        a = read
        b = read
        {:ok, [l,r]} = :io.fread('','~d~d')
        
        p = Enum.zip(a, b)

        n = (r-l)*1000        
        xi = Enum.map(0..n, &(l + &1*0.001))
        
        area = sum(p, xi, &(0.001*f(&1, &2)))
        volume = sum(p, xi, &area(&1, &2))
        
        IO.puts area
        IO.puts volume
    end
end

Solution.main
