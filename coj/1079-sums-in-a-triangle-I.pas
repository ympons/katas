 var
     m : array[0..100,0..100] of integer;
     n , i, p, j , k, sum: integer;

  function max(a, b : integer): integer;
  begin
       if (a > b) then max := a
                  else max := b;
  end;

  begin
       readln(n);
       for i:=1 to n do
       begin
            readln(p);
            fillchar(m, sizeof(m),0);
            sum := 0;
            for k:=1 to p do
               for j:=1 to k do
               begin
                    read(m[k][j]);
                    m[k][j] := max(m[k][j] + m[k-1][j], m[k][j] + m[k-1][j-1]);
                    sum := max(sum, m[k][j]);
               end;
            writeln(sum);
       end;
  end.
