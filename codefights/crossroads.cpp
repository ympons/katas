bool crossroads(std::vector<int> road1, std::vector<int> road2,
                int crossTime) {
  int j = 0;
  for (int i = 0; i < road1.size(); i++) {
    while (j < road2.size() && road1[i] + crossTime > road2[j] ) {
      if (road2[j] + crossTime > road1[i]) {
        return true;
      }
      j++;
    }
    if (j < road2.size() && road1[i] + crossTime > road2[j]) {
      return true;
    }
  }
  return false;
}


/*
Consider an Uber-City where the only passenger vehicles on the roads are self-driving Uber cars. Since all the cars are self-driving, you could synchronize the car speeds in such a way that traffic lights would be obsolete.

As a first step you need to implement an algorithm that would check whether removing traffic lights at a given intersection of two one-way roads will lead to a car crash.

For each road you know when a car will approach the crossroads. You also know how long it takes to cross the crossroads.

Example

For road1 = [1, 5, 6, 7], road2 = [3, 10] and crossTime = 2, the output should be
crossroads(road1, road2, crossTime) = false.

The cars will cross the crossroads in the following order without colliding: road1[0], road2[0], road1[1], road1[2], road1[3], road2[1].

For road1 = [2], road2 = [3] and crossTime = 2, the output should be
crossroads(road1, road2, crossTime) = true.

The car on the first road will not make it when the car from the second road appears since crossing the intersection takes 4 seconds.

Input/Output

[time limit] 500ms (cpp)
[input] array.integer road1

A strictly increasing array of positive integers which define the times (in seconds) when a new car approaches the crossroads by the first road.

Constraints:
1 ≤ road1.length ≤ 10,
1 ≤ road1[i] ≤ 40.

[input] array.integer road2

A strictly increasing array of positive integers which define the times (in seconds) when a new car approaches the crossroads by the second road.

Constraints:
0 ≤ road2.length ≤ 10,
1 ≤ road2[i] ≤ 40.

[input] integer crossTime

The number of seconds it takes for a car to cross the crossroads.

Constraints:
1 ≤ crossTime ≤ 10.

[output] boolean

true if removing the traffic lights will cause a collision, false otherwise.
*/