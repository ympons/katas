object Solution {
    
    def fibonacci(x:Int):Int = {
         if (x <= 3) 1 else fibonacci(x-1) + fibonacci(x-2)
    }

    def main(args: Array[String]) {
         println(fibonacci(readInt()))
    }
}

