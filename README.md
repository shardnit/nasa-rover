NASA Rover
==========
This project consists of a `Rover` struct which has `Init()` and `Move()` methods. We assume that the origin of rectangular 
plateau is `(0, 0)`. 

# Move Logic
- Depending on the current direction of the rover, given `L` or `R` command it will spin left or right respectively.
- If `M` command is provided, it will move one grid forward (again depends on the current direction).
- If the rover is already at the boundary, then depending on the current direction and command it will not cross the boundary.

# Run Tests
- To run tests: `go test .` or `go test -v .` (for a verbose output).

# Run project
- First build the project using `go build`. This would result in a binary named `nasa-rover` in current dir.
- `./nasa-rover <input-file-name>`. This project comes with a sample input file.
